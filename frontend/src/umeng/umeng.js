export default {
	push(category,action,label,value,nodeid) {
		if(window._czc){
			window._czc.push(["_trackEvent",category,action,label,value,nodeid])
		}else {
			console.error('error')
		}
	}
}
