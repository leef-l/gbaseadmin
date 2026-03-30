package playapi

import "gbaseadmin/app/play/internal/service"

func init() {
	service.RegisterPlayapiAuth(&sAuth{})
	service.RegisterPlayapiMember(&sMember{})
	service.RegisterPlayapiCoach(&sPlayapiCoach{})
	service.RegisterPlayapiGoods(&sPlayapiGoods{})
	service.RegisterPlayapiOrder(&sPlayapiOrder{})
	service.RegisterPlayapiPayment(&sPlayapiPayment{})
	service.RegisterPlayapiReview(&sPlayapiReview{})
	service.RegisterPlayapiRecharge(&sRecharge{})
	service.RegisterPlayapiCoupon(&sCoupon{})
	service.RegisterPlayapiActivity(&sActivity{})
	service.RegisterPlayapiSearch(&sSearch{})
	service.RegisterPlayapiMessage(&sMessage{})
}
