import { TX_TYPE } from "../constant/Constant";


export class TxHelper {

    static getFromAndToAddressFromMsg(msgs){

        let res = {
            from : '--',
            to : '--'
        };
        if(!msgs || !msgs.msg) return res;
        const {type, msg} = msgs;
        switch (type){
            case TX_TYPE.send:
                res.from = msg.fromaddress;
                res.to = msg.toaddress;
                break;
            case TX_TYPE.define_service:
                res.from = msg.author;
                break;
            case TX_TYPE.bind_service:
                res.from = msg.provider;
                break;
            case TX_TYPE.update_service_binding:
                res.from = msg.provider;
                break;
            case TX_TYPE.disable_service_binding:
                res.from = msg.provider;
                break;
            case TX_TYPE.enable_service_binding:
                res.from = msg.provider;
                break;
            case TX_TYPE.refund_service_deposit:
                res.from = msg.provider;
                break;
            case TX_TYPE.call_service:
                res.from = msg.consumer;
                if(msg.providers.length === 0){
                    res.to = '--'
                } else {
                    if(msg.providers.length > 1){
                        res.to = msg.providers;
                    } else {
                        res.to = msg.providers[0]
                    }
                }
                break;
            case TX_TYPE.respond_service:
                res.from = msg.provider;
                if(msg.ex && msg.ex.consumer){
                    res.to = msg.ex.consumer;
                }
                break;
            case TX_TYPE.pause_request_context:
                res.from = msg.consumer;
                break;
            case TX_TYPE.start_request_context:
                res.from = msg.consumer;
                break;
            case TX_TYPE.kill_request_context:
                res.from = msg.consumer;
                break;
            case TX_TYPE.update_request_context:
                res.from = msg.consumer;
                break;
            case TX_TYPE.create_record:
                res.from = msg.creator;
                break;
            case TX_TYPE.burn_nft:
                res.from = msg.sender;
                break;
            case TX_TYPE.transfer_nft:
                res.from = msg.sender;
                res.to = msg.recipient;
                break;
            case TX_TYPE.edit_nft:
                res.from = msg.sender;
                break;
            case TX_TYPE.issue_denom:
                res.from = msg.sender;
                break;
            case TX_TYPE.mint_nft:
                res.from = msg.sender;
                res.to = msg.recipient;
                break;
        }
        return res;
    }

    static getContextId(msgs, events){
        if(!msgs) return;
        let requestContextId = '';
        const {type, msg} = msgs;
        switch (type){
            case TX_TYPE.define_service:
                requestContextId = '--';
                break;
            case TX_TYPE.bind_service:
                requestContextId = '--';
                break;
            case TX_TYPE.update_service_binding:
                requestContextId = '--';
                break;
            case TX_TYPE.disable_service_binding:
                requestContextId = '--';
                break;
            case TX_TYPE.enable_service_binding:
                requestContextId = '--';
                break;
            case TX_TYPE.refund_service_deposit:
                requestContextId = '--';
                break;
            case TX_TYPE.call_service:
                if(events && Array.isArray(events)){
                    for(let e of events){
                        if(e && e.attributes && Array.isArray(e.attributes)){
                            for(let a of e.attributes){
                                if(a.key === 'request_context_id'){
                                    requestContextId = a.value;
                                    break;
                                }
                            }
                        }

                    }
                }
                break;
            case TX_TYPE.respond_service:
                requestContextId = msg.ex ? msg.ex.request_context_id : '';
                break;
            case TX_TYPE.pause_request_context:
                requestContextId = msg.request_context_id;
                break;
            case TX_TYPE.start_request_context:
                requestContextId = msg.request_context_id;
                break;
            case TX_TYPE.kill_request_context:
                requestContextId = msg.request_context_id;
                break;
            case TX_TYPE.update_request_context:
                requestContextId = msg.request_context_id;
                break;
        }
        return requestContextId;
    }

}