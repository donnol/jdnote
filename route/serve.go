package route

import "fmt"

// StartServer 开启服务
func StartServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	if err := defaultRouter.Run(addr); err != nil {
		return err
	}

	return nil
}
