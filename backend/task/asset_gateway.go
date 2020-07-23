package task

//type UpdateAssetGateways struct{}
//
//func (task UpdateAssetGateways) Name() string {
//	return "update_asset_gateways"
//}
//
//func (task UpdateAssetGateways) Start() {
//	timeInterval := conf.Get().Server.CronTimeAssetGateways
//	taskName := task.Name()
//
//	utils.RunTimer(timeInterval, utils.Sec, func() {
//		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
//			logger.Error(err.Error())
//		}
//	})
//}
//
//func (task UpdateAssetGateways) DoTask(fn func(string) chan bool) error {
//	stop := fn(task.Name())
//	defer HeartQuit(stop)
//	assetGateways, err := document.AssetGateways{}.GetAllAssetGateways()
//	if err != nil {
//		return err
//	}
//
//	assetService := service.Get(service.Asset).(*service.AssetsService)
//	if err := assetService.UpdateAssetGateway(assetGateways); err != nil {
//		return err
//	}
//
//	return nil
//}
