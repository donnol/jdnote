-- 时间更新触发器
CREATE FUNCTION update_updated_at() RETURNS trigger AS $update_updated_at$
    BEGIN
        NEW.updated_at := clock_timestamp();
        RETURN NEW;
    END;
$update_updated_at$ LANGUAGE plpgsql;
