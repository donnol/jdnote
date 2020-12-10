package influx_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/donnol/jdnote/utils/store/influx"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestInflux(t *testing.T) {
	ctx := context.Background()

	client := influx.Open(influx.Option{
		Host:  "http://localhost:8086",
		Token: "zkKjAsnZ8_5-e6kAWytj-li_LZvusdfCGgaXmxZiktzUcJj5yueasLjKVUyhYgKkDeYKMVP8cMsPIMzi5rY1RA==",
	}, nil)

	hc, err := client.Health(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hc: %+v\n", hc)
	if hc != nil {
		t.Logf("%v\n", *hc.Message) // ready for queries and writes
	}

	bs := influx.BucketSetting{
		OrgName:    "jdorg",
		BucketName: "jdbucket",
	}

	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(bs.OrgName, bs.BucketName)
	// Create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45.0},
		time.Now())
	// write point immediately
	if err := writeAPI.WritePoint(ctx, p); err != nil {
		t.Fatal(err)
	}

	// Create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45.0).
		SetTime(time.Now())
	if err := writeAPI.WritePoint(ctx, p); err != nil {
		t.Fatal(err)
	}

	// Or write directly line protocol
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	if err := writeAPI.WriteRecord(ctx, line); err != nil {
		t.Fatal(err)
	}

	// Get query client
	queryAPI := client.QueryAPI(bs.OrgName)
	// Get parser flux query result
	result, err := queryAPI.Query(ctx, `from(bucket:"`+bs.BucketName+`")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			// Observe when there is new grouping key producing new table
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// read result
			fmt.Printf("row: %s\n", result.Record().String())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	}

	// Ensures background processes finishes
	client.Close()
}
