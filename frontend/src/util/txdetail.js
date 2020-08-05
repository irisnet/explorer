import Constant from "../constant/Constant"
import numberMoveDecimal from "move-decimal-point"
import Tools from "../util/Tools"
import bigNumber from 'bignumber.js'
import moment from 'moment';

export default class formatMsgsAndTags {
    static formatAmount(amount){
        amount = amount.toString()
        if(amount.includes('.') || amount.includes(',')){
            return amount
        }else {
            return new bigNumber(amount).toFormat()
        }
    }

    static switchTxType(dataTx, ctx){
        let message;
        switch (dataTx.type){
            case Constant.TxType.TRANSFER :
                return message = formatMsgsAndTags.txTypeTransfer(dataTx, Constant.TxType.TRANSFER);
                break;
            case Constant.TxType.BURN :
                return message = formatMsgsAndTags.txTypeBurn(dataTx, Constant.TxType.BURN);
                break;
            case Constant.TxType.SETMEMOREGEXP :
                return message = formatMsgsAndTags.txTypeSetMemoRegexp(dataTx, Constant.TxType.SETMEMOREGEXP);
                break;
            case Constant.TxType.CREATEVALIDATOR :
                return message = formatMsgsAndTags.txTypeCreateValidator(dataTx, Constant.TxType.CREATEVALIDATOR);
                break;
            case Constant.TxType.EDITVALIDATOR :
                return message = formatMsgsAndTags.txTypeEditValidator(dataTx, Constant.TxType.EDITVALIDATOR);
                break;
            case Constant.TxType.UNJAIL :
                return message = formatMsgsAndTags.txTypeUnjail(dataTx, Constant.TxType.UNJAIL);
                break;
            case Constant.TxType.DELEGATE :
                return message = formatMsgsAndTags.txTypeDelegate(dataTx, Constant.TxType.DELEGATE);
                break;
            case Constant.TxType.BEGINREDELEGATE :
                return message = formatMsgsAndTags.txTypeBeginRedelegate(dataTx, Constant.TxType.BEGINREDELEGATE);
                break;
            case Constant.TxType.SETWITHDRAWADDRESS :
                return message = formatMsgsAndTags.txTypeSetWithdrawAddress(dataTx, Constant.TxType.SETWITHDRAWADDRESS);
                break;
            case Constant.TxType.BEGINUNBONDING :
                return message = formatMsgsAndTags.txTypeBeginUnbonding(dataTx, Constant.TxType.BEGINUNBONDING);
                break;
            case Constant.TxType.WITHDRAWDELEGATORREWARD :
                return message = formatMsgsAndTags.txTypeWithdrawDelegatorReward(dataTx, Constant.TxType.WITHDRAWDELEGATORREWARD);
                break;
            case Constant.TxType.WITHDRAWDELEGATORREWARDSALL :
                return message = formatMsgsAndTags.txTypeWithdrawDelegatorRewardsAll(dataTx, Constant.TxType.WITHDRAWDELEGATORREWARDSALL);
                break;
            case Constant.TxType.WITHDRAWVALIDATORREWARDSALL :
                return message = formatMsgsAndTags.txTypeWithdrawValidatorRewardsAll(dataTx, Constant.TxType.WITHDRAWVALIDATORREWARDSALL);
                break;
            case Constant.TxType.SUBMITPROPOSAL :
                return message = formatMsgsAndTags.txTypeSubmitProposal(dataTx, Constant.TxType.SUBMITPROPOSAL);
                break;
            case Constant.TxType.DEPOSIT :
                return message = formatMsgsAndTags.txTypeDeposit(dataTx, Constant.TxType.DEPOSIT);
                break;
            case Constant.TxType.VOTE :
                return message = formatMsgsAndTags.txTypeVote(dataTx, Constant.TxType.VOTE);
                break;
            case Constant.TxType.ISSUETOKEN :
                return message = formatMsgsAndTags.txTypeIssueToken(dataTx, Constant.TxType.ISSUETOKEN);
                break;
            case Constant.TxType.EDITTOKEN :
                return message = formatMsgsAndTags.txTypeEditToken(dataTx, Constant.TxType.EDITTOKEN);
                break;
            case Constant.TxType.MINTTOKEN :
                return message = formatMsgsAndTags.txTypeMintToken(dataTx, Constant.TxType.MINTTOKEN);
                break;
            case Constant.TxType.TRANSFERTOKENOWNER :
                return message = formatMsgsAndTags.txTypeTransferTokenOwner(dataTx, Constant.TxType.TRANSFERTOKENOWNER);
                break;
            case Constant.TxType.CREATEGATEWAY :
                return message = formatMsgsAndTags.txTypeCreateGateway(dataTx, Constant.TxType.CREATEGATEWAY);
                break;
            case Constant.TxType.EDITGATEWAY :
                return message = formatMsgsAndTags.txTypeEditGateway(dataTx, Constant.TxType.EDITGATEWAY);
                break;
            case Constant.TxType.TRANSFERGATEWAYOWNER :
                return message = formatMsgsAndTags.txTypeTransferGatewayOwner(dataTx, Constant.TxType.TRANSFERGATEWAYOWNER);
                break;
            case Constant.TxType.REQUESTRAND :
                return message = formatMsgsAndTags.txTypeRequestRand(dataTx, Constant.TxType.REQUESTRAND);
                break;
            case Constant.TxType.ADDPROFILER :
                return message = formatMsgsAndTags.txTypeAddProfiler(dataTx, Constant.TxType.ADDPROFILER);
                break;
            case Constant.TxType.ADDTRUSTEE :
                return message = formatMsgsAndTags.txTypeAddTrustee(dataTx, Constant.TxType.ADDTRUSTEE);
                break;
            case Constant.TxType.DELETEPROFIKER :
                return message = formatMsgsAndTags.txTypeDeleteProfiler(dataTx, Constant.TxType.DELETEPROFIKER);
                break;
            case Constant.TxType.DELETETRUSTEE :
                return message = formatMsgsAndTags.txTypeDeleteTrustee(dataTx, Constant.TxType.DELETETRUSTEE);
                break;
            case Constant.TxType.CLAIMHTLC :
                return message = formatMsgsAndTags.txTypeClaimHTLC(dataTx, Constant.TxType.CLAIMHTLC);
                break;
            case Constant.TxType.CREATEHTLC :
                return message = formatMsgsAndTags.txTypeCreateHTLC(dataTx, Constant.TxType.CREATEHTLC);
                break;
            case Constant.TxType.REFUNDHTLC :
                return message = formatMsgsAndTags.txTypeRefundHTLC(dataTx, Constant.TxType.REFUNDHTLC);
                break;
            case Constant.TxType.DEFINE_SERVICE :
                return formatMsgsAndTags.txTypeDefineService(dataTx);
            case Constant.TxType.BIND_SERVICE :
                return formatMsgsAndTags.txTypeBindService(dataTx);
            case Constant.TxType.UPDATE_SERVICE_BINDING :
                return formatMsgsAndTags.txTypeUpdateServiceBinding(dataTx);
            case Constant.TxType.SET_WITHDRAW_FEES_ADDRESS :
                return formatMsgsAndTags.txTypeSetWithdrawFeesAddress(dataTx);
            case Constant.TxType.DISABLE_SERVICE_BINDING :
                return formatMsgsAndTags.txTypeDisableServiceBinding(dataTx);
            case Constant.TxType.ENABLE_SERVICE_BINDING :
                return formatMsgsAndTags.txTypeEnableServiceBinding(dataTx);
            case Constant.TxType.REFUND_SERVICE_DEPOSIT :
                return formatMsgsAndTags.txTypeRefundServiceDeposit(dataTx);
            case Constant.TxType.CALL_SERVICE :
                return formatMsgsAndTags.txTypeCallService(dataTx);

            case Constant.TxType.RESPOND_SERVICE :
                return formatMsgsAndTags.txTypeRespondService(dataTx);

            case Constant.TxType.PAUSE_REQUEST_CONTEXT :
                return formatMsgsAndTags.txTypePauseRequestContext(dataTx);
            case Constant.TxType.START_REQUEST_CONTEXT :
                return formatMsgsAndTags.txTypeStartRequestContext(dataTx);
            case Constant.TxType.KILL_REQUEST_CONTEXT :
                return formatMsgsAndTags.txTypeKillRequestContext(dataTx);
            case Constant.TxType.UPDATE_REQUEST_CONTEXT :
                return formatMsgsAndTags.txTypeUpdateRequestContext(dataTx);
            case Constant.TxType.WITHDRAW_EARNED_FEES :
                return formatMsgsAndTags.txTypeWithdrawEarnedFees(dataTx);

            case Constant.TxType.FundCommunityPool :
                return formatMsgsAndTags.txTypeFundCommunityPool(dataTx);
            case Constant.TxType.ISSUE_DENOM :
                return formatMsgsAndTags.txTypeIssueDenom(dataTx);
            case Constant.TxType.EDIT_NFT :
                return formatMsgsAndTags.txTypeEditNft(dataTx);
            case Constant.TxType.TRANSFER_NFT :
                return formatMsgsAndTags.txTypeTransferNft(dataTx);
            case Constant.TxType.MINT_NFT :
                return formatMsgsAndTags.txTypeMintNft(dataTx);
            case Constant.TxType.BURN_NFT :
                return formatMsgsAndTags.txTypeBurnNft(dataTx);
            case Constant.TxType.CREATE_FEED :
                return formatMsgsAndTags.txTypeCreateFeed(dataTx);
            case Constant.TxType.START_FEED :
                return formatMsgsAndTags.txTypeStartFeed(dataTx);
            case Constant.TxType.PAUSE_FEED :
                return formatMsgsAndTags.txTypePauseFeed(dataTx);
            case Constant.TxType.EDIT_FEED :
                return formatMsgsAndTags.txTypeEditFeed(dataTx);
            case Constant.TxType.ADDLIQUIDITY :
                return formatMsgsAndTags.txTypeAddLiquidity(dataTx, ctx);
            case Constant.TxType.REMOVELIQUIDITY :
                return formatMsgsAndTags.txTypeRemoveLiquidity(dataTx, ctx);
            case Constant.TxType.SWAPORDER :
                return formatMsgsAndTags.txTypeSwapOrder(dataTx, ctx);




        }

    }

    static txTypeDefineService(tx){
        const msg = {
            'Type :' : [],
            'Name :' : [],
            'Description :' : [],
            'Tags :' : [],
            'Author :' : [],
            'Author Description :' : [],
            'Schemas :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Name :'].push(m.msg.name || '--');
                msg['Description :'].push(m.msg.description || '--');
                msg['Tags :'].push(m.msg.tags);
                msg['Author :'].push(m.msg.author || '--');
                msg['Author Description :'].push(m.msg.authorDescription || '--');
                msg['Schemas :'].push(m.msg.schemas);
            })
        }
        return msg;
    }

    static txTypeBindService(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Pricing :' : [],
            'QoS :' : [],
            'Deposit :' : [],
            'Provider :' : [],
            'Owner :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Service Name :'].push(m.msg.service_name || '--');
                msg['Provider :'].push(m.msg.provider || '--');
                if(m.msg.deposit && Array.isArray(m.msg.deposit) && m.msg.deposit.length > 0){
                    msg['Deposit :'].push(`${m.msg.deposit[0].amount} ${m.msg.deposit[0].denom}`);
                } else {
                    msg['Deposit :'].push('--')
                }

                msg['Pricing :'].push(m.msg.pricing ? JSON.parse(m.msg.pricing).price : '--');
                msg['QoS :'].push(`${m.msg.qos} blocks`);
                msg['Owner :'].push(m.msg.owner);
            })
        }
        return msg;
    }

    static txTypeUpdateServiceBinding(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Pricing :' : [],
            'QoS :' : [],
            'Deposit :' : [],
            'Provider :' : [],
            'Owner :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Service Name :'].push(m.msg.service_name || '--');
                msg['Provider :'].push(m.msg.provider);
                if(m.msg.deposit && Array.isArray(m.msg.deposit) && m.msg.deposit.length > 0){
                    msg['Deposit :'].push(`${m.msg.deposit[0].amount} ${m.msg.deposit[0].denom}`);
                } else {
                    msg['Deposit :'].push('--')
                }

                msg['Pricing :'].push(m.msg.pricing ? JSON.parse(m.msg.pricing).price : '--');
                msg['QoS :'].push(`${m.msg.qos} blocks`);
                msg['Owner :'].push(m.msg.owner);
            })
        }
        return msg;
    }

    static txTypeSetWithdrawFeesAddress(tx){
        const msg = {
            'Type :' : [],
            'Owner :' : [],
            'Withdraw Address :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Owner :'].push(m.msg.owner);
                msg['Withdraw Address :'].push(m.msg.withdraw_address);
            })
        }
        return msg;
    }

    static txTypeDisableServiceBinding(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Provider :' : [],
            'Owner :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Service Name :'].push(m.msg.service_name || '--');
                msg['Provider :'].push(m.msg.provider);
                msg['Owner :'].push(m.msg.owner);
            })
        }
        return msg;
    }

    static txTypeEnableServiceBinding(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Deposit :' : [],
            'Provider :' : [],
            'Owner :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Service Name :'].push(m.msg.service_name || '--');
                msg['Provider :'].push(m.msg.provider);
                if(m.msg.deposit && Array.isArray(m.msg.deposit) && m.msg.deposit.length > 0){
                    msg['Deposit :'].push(`${m.msg.deposit[0].amount} ${m.msg.deposit[0].denom}`);
                } else {
                    msg['Deposit :'].push('--')
                }
                msg['Owner :'].push(m.msg.owner);
            })
        }
        return msg;
    }

    static txTypeRefundServiceDeposit(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Provider :' : [],
            'Owner :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Service Name :'].push(m.msg.service_name || '--');
                msg['Provider :'].push(m.msg.provider);
                msg['Owner :'].push(m.msg.owner);
            })
        }
        return msg;
    }


    static txTypeCallService(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Request Context ID :' : [],
            'Consumer :' : [],
            'Input :' : [],
            'Providers :' : [],
            'Repeated :' : [],
            'Repeated Frequency :' : [],
            'Repeated Total :' : [],
            'Service Fee Cap :' : [],
            'Super Mode :' : [],
            'Time Out :' : [],


        };
        msg['Type :'].push(tx.type);
        if(tx && tx.events && Array.isArray(tx.events) && tx.events.length > 0){
            for(let e of tx.events){
                if(e.attributes.request_context_id){
                    msg['Request Context ID :'].push(e.attributes.request_context_id);
                }
            }
        }
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                if(m.msg.ex && m.msg.ex.service_name){
                    msg['Service Name :'].push(m.msg.ex.service_name);
                } else {
                    msg['Service Name :'].push('--');
                }
                if(m.msg.providers && Array.isArray(m.msg.providers) && m.msg.providers.length > 0){
                    m.msg.providers.forEach((p) =>{
                        msg['Providers :'].push(p);
                    })
                } else {
                    msg['Providers :'].push('--');
                }
                msg['Consumer :'].push(m.msg.consumer);
                msg['Input :'].push(m.msg.input);
                if(m.msg.service_fee_cap && Array.isArray(m.msg.service_fee_cap) && m.msg.service_fee_cap.length > 0){
                    msg['Service Fee Cap :'].push(`${m.msg.service_fee_cap[0].amount} ${m.msg.service_fee_cap[0].denom}`);
                } else {
                    msg['Service Fee Cap :'].push('--')
                }
                msg['Time Out :'].push(m.msg.timeout);
                msg['Super Mode :'].push(m.msg.super_mode);
                msg['Repeated :'].push(m.msg.repeated);
                msg['Repeated Frequency :'].push(m.msg.repeated_frequency);
                msg['Repeated Total :'].push(m.msg.repeated_total);
            })
        }
        return msg;
    }

    static txTypeRespondService(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Request Context ID :' : [],
            'Provider :' : [],
            'Result :' : [],
            'Output :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                if(m.msg.ex && m.msg.ex.service_name){
                    msg['Service Name :'].push(m.msg.ex.service_name);
                } else {
                    msg['Service Name :'].push('--');
                }
                if(m.msg.ex && m.msg.ex.request_context_id){
                    msg['Request Context ID :'].push(m.msg.ex.request_context_id);
                } else {
                    msg['Request Context ID :'].push('--');
                }

                msg['Provider :'].push(m.msg.provider);
                msg['Result :'].push(m.msg.result);
                msg['Output :'].push(m.msg.output);
            })
        }
        return msg;
    }

    static txTypePauseRequestContext(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Request Context ID :' : [],
            'Consumer :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                if(m.msg.ex && m.msg.ex.service_name){
                    msg['Service Name :'].push(m.msg.ex.service_name);
                } else {
                    msg['Service Name :'].push('--');
                }
                msg['Request Context ID :'].push(m.msg.request_context_id);
                msg['Consumer :'].push(m.msg.consumer);
            })
        }
        return msg;
    }

    static txTypeStartRequestContext(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Request Context ID :' : [],
            'Consumer :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                if(m.msg.ex && m.msg.ex.service_name){
                    msg['Service Name :'].push(m.msg.ex.service_name);
                } else {
                    msg['Service Name :'].push('--');
                }
                msg['Request Context ID :'].push(m.msg.request_context_id);
                msg['Consumer :'].push(m.msg.consumer);
            })
        }
        return msg;
    }

    static txTypeKillRequestContext(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Request Context ID :' : [],
            'Consumer :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                if(m.msg.ex && m.msg.ex.service_name){
                    msg['Service Name :'].push(m.msg.ex.service_name);
                } else {
                    msg['Service Name :'].push('--');
                }
                msg['Request Context ID :'].push(m.msg.request_context_id);
                msg['Consumer :'].push(m.msg.consumer);
            })
        }
        return msg;
    }

    static txTypeUpdateRequestContext(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Request Context ID :' : [],
            'Consumer :' : [],
            'Providers :' : [],
            'Repeated Frequency :' : [],
            'Repeated Total :' : [],
            'Service Fee Cap :' : [],
            'Time Out :' : [],


        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                if(m.msg.ex && m.msg.ex.service_name){
                    msg['Service Name :'].push(m.msg.ex.service_name);
                } else {
                    msg['Service Name :'].push('--');
                }
                msg['Request Context ID :'].push(m.msg.request_context_id);
                msg['Consumer :'].push(m.msg.consumer);
                if(m.msg.providers && Array.isArray(m.msg.providers) && m.msg.providers.length > 0){
                    m.msg.providers.forEach((p) =>{
                        msg['Providers :'].push(p);
                    })
                } else {
                    msg['Providers :'].push('--');
                }

                if(m.msg.service_fee_cap && Array.isArray(m.msg.service_fee_cap) && m.msg.service_fee_cap.length > 0){
                    msg['Service Fee Cap :'].push(`${m.msg.service_fee_cap[0].amount} ${m.msg.service_fee_cap[0].denom}`);
                } else {
                    msg['Service Fee Cap :'].push('--')
                }
                msg['Time Out :'].push(m.msg.timeout);
                msg['Repeated Frequency :'].push(m.msg.repeated_frequency);
                msg['Repeated Total :'].push(m.msg.repeated_total);


            })
        }
        return msg;
    }

    static txTypeWithdrawEarnedFees(tx){
        const msg = {
            'Type :' : [],
            'Provider :' : [],
            'Owner :' : [],

        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Provider :'].push(m.msg.provider);
                msg['Owner :'].push(m.msg.owner);
            })
        }
        return msg;
    }

    static txTypeFundCommunityPool(tx){
        const msg = {
            'Type :' : [],
            'Amount :' : [],
            'Depositor :' : [],

        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                if(m.msg.amount && Array.isArray(m.msg.amount) && m.msg.amount.length > 0){
                    msg['Amount :'].push(`${m.msg.amount[0].amount} ${m.msg.amount[0].denom}`);
                }
                msg['Depositor :'].push(m.msg.depositor);
            })
        }
        return msg;
    }

    static txTypeIssueDenom(tx){
        const msg = {
            'Type :' : [],
            'From :' : [],
            'Denom :' : [],
            'Schema :' : [],

        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['From :'].push(m.msg.sender);
                msg['Denom :'].push(m.msg.denom);
                msg['Schema :'].push(m.msg.schema);
            })
        }
        return msg;
    }

    static txTypeEditNft(tx){
        const msg = {
            'Type :' : [],
            'From :' : [],
            'Denom :' : [],
            'NFT Id :' : [],
            'NFT Uri :' : [],
            'NFT Data :' : [],

        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['From :'].push(m.msg.sender);
                msg['NFT Id :'].push(m.msg.id);
                msg['Denom :'].push(m.msg.denom);
                msg['NFT Uri :'].push(m.msg.token_uri);
                msg['NFT Data :'].push(m.msg.token_data);
            })
        }
        return msg;
    }

    static txTypeTransferNft(tx){
        const msg = {
            'Type :' : [],
            'From :' : [],
            'Recipient :' : [],
            'Denom :' : [],
            'NFT Id :' : [],
            'NFT Uri :' : [],
            'NFT Data :' : [],

        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['From :'].push(m.msg.sender);
                msg['Recipient :'].push(m.msg.recipient);
                msg['NFT Id :'].push(m.msg.id);
                msg['Denom :'].push(m.msg.denom);
                msg['NFT Uri :'].push(m.msg.token_uri);
                msg['NFT Data :'].push(m.msg.token_data);
            })
        }
        return msg;
    }



    static txTypeMintNft(tx){
        const msg = {
            'Type :' : [],
            'From :' : [],
            'Recipient :' : [],
            'Denom :' : [],
            'NFT Id :' : [],
            'NFT Uri :' : [],
            'NFT Data :' : [],

        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['From :'].push(m.msg.sender);
                msg['Recipient :'].push(m.msg.recipient);
                msg['NFT Id :'].push(m.msg.id);
                msg['Denom :'].push(m.msg.denom);
                msg['NFT Uri :'].push(m.msg.token_uri);
                msg['NFT Data :'].push(m.msg.token_data);
            })
        }
        return msg;
    }

    static txTypeBurnNft(tx){
        const msg = {
            'Type :' : [],
            'From :' : [],
            'Denom :' : [],
            'NFT Id :' : [],

        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['From :'].push(m.msg.sender);
                msg['NFT Id :'].push(m.msg.id);
                msg['Denom :'].push(m.msg.denom);
            })
        }
        return msg;
    }

    static txTypeCreateFeed(tx){
        const msg = {
            'Type :' : [],
            'Service Name :' : [],
            'Feed Name :' : [],
            'Description :' : [],
            'Latest History :' : [],
            'Creator :' : [],
            'Providers :' : [],
            'Input :' : [],
            'Service Fee Cap :' : [],
            'Aggregate Func :' : [],
            'Value Json Path :' : [],
            'Repeated Frequency :' : [],
            'Response Threshold :' : [],
            'Timeout :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Feed Name :'].push(m.msg.feed_name);
                msg['Latest History :'].push(m.msg.latest_history);
                msg['Description :'].push(m.msg.description);
                msg['Creator :'].push(m.msg.creator);
                msg['Service Name :'].push(m.msg.service_name);
                msg['Input :'].push(m.msg.input);
                msg['Timeout :'].push(m.msg.timeout);
                msg['Repeated Frequency :'].push(m.msg.repeated_frequency);
                msg['Aggregate Func :'].push(m.msg.aggregate_func);
                msg['Value Json Path :'].push(m.msg.value_json_path);
                msg['Response Threshold :'].push(m.msg.response_threshold);
                if(m.msg.providers && Array.isArray(m.msg.providers) && m.msg.providers.length > 0){
                    m.msg.providers.forEach((p) =>{
                        msg['Providers :'].push(p);
                    })
                } else {
                    msg['Providers :'].push('--');
                }

                if(m.msg.service_fee_cap && Array.isArray(m.msg.service_fee_cap) && m.msg.service_fee_cap.length > 0){
                    msg['Service Fee Cap :'].push(`${m.msg.service_fee_cap[0].amount} ${m.msg.service_fee_cap[0].denom}`);
                } else {
                    msg['Service Fee Cap :'].push('--')
                }
            })
        }
        return msg;
    }

    static txTypeStartFeed(tx){
        const msg = {
            'Type :' : [],
            'Feed Name :' : [],
            'Creator :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Feed Name :'].push(m.msg.feed_name);
                msg['Creator :'].push(m.msg.creator);
            })
        }
        return msg;
    }

    static txTypePauseFeed(tx){
        const msg = {
            'Type :' : [],
            'Feed Name :' : [],
            'Creator :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Feed Name :'].push(m.msg.feed_name);
                msg['Creator :'].push(m.msg.creator);
            })
        }
        return msg;
    }

    static txTypeEditFeed(tx){
        const msg = {
            'Type :' : [],
            'Feed Name :' : [],
            'Description :' : [],
            'Latest History :' : [],
            'Creator :' : [],
            'Providers :' : [],
            'Service Fee Cap :' : [],
            'Repeated Frequency :' : [],
            'Response Threshold :' : [],
            'Timeout :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Feed Name :'].push(m.msg.feed_name);
                msg['Latest History :'].push(m.msg.latest_history);
                msg['Description :'].push(m.msg.description);
                msg['Creator :'].push(m.msg.creator);
                msg['Timeout :'].push(m.msg.timeout);
                msg['Repeated Frequency :'].push(m.msg.repeated_frequency);
                msg['Response Threshold :'].push(m.msg.response_threshold);
                if(m.msg.providers && Array.isArray(m.msg.providers) && m.msg.providers.length > 0){
                    m.msg.providers.forEach((p) =>{
                        msg['Providers :'].push(p);
                    })
                } else {
                    msg['Providers :'].push('--');
                }

                if(m.msg.service_fee_cap && Array.isArray(m.msg.service_fee_cap) && m.msg.service_fee_cap.length > 0){
                    msg['Service Fee Cap :'].push(`${m.msg.service_fee_cap[0].amount} ${m.msg.service_fee_cap[0].denom}`);
                } else {
                    msg['Service Fee Cap :'].push('--')
                }
            })
        }
        return msg;
    }


    static txTypeAddLiquidity(tx, ctx){
        const msg = {
            'Type :' : [],
            'Sender :' : [],
            'Exact Iris Amt :' : [],
            'Max Token :' : [],
            'Min Liquidity :' : [],
            'Deadline :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Sender :'].push(m.msg.sender);
                const amt = Tools.numberMoveDecimal(m.msg.exact_iris_amt);
                const denom = ctx.$store.state.nativeToken;
                msg['Exact Iris Amt :'].push(`${formatMsgsAndTags.formatAmount(amt)} ${denom}`);
                if(m.msg.max_token.denom && m.msg.max_token.amount){
                    msg['Max Token :'].push(`${m.msg.max_token.amount} ${m.msg.max_token.denom}`);
                }
                msg['Min Liquidity :'].push(m.msg.min_liquidity);
                msg['Deadline :'].push(Tools.getDisplayDate(m.msg.deadline*1000));
            })
        }
        return msg;
    }

    static txTypeRemoveLiquidity(tx, ctx){
        const msg = {
            'Type :' : [],
            'Sender :' : [],
            'Withdraw Liquidity :' : [],
            'Min Iris Amt :' : [],
            'Min Token :' : [],
            'Deadline :' : [],
        };
        msg['Type :'].push(tx.type);
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{
                msg['Sender :'].push(m.msg.sender);
                if(m.msg.withdraw_liquidity.denom && m.msg.withdraw_liquidity.amount){

                    msg['Withdraw Liquidity :'].push(`${m.msg.withdraw_liquidity.amount} ${m.msg.withdraw_liquidity.denom}`);
                }
                const amt = Tools.numberMoveDecimal(m.msg.min_iris_amt);
                const denom = ctx.$store.state.nativeToken;
                msg['Min Iris Amt :'].push(`${formatMsgsAndTags.formatAmount(amt)} ${denom}`);
                msg['Min Token :'].push(m.msg.min_token);
                msg['Deadline :'].push(Tools.getDisplayDate(m.msg.deadline*1000));
            })
        }
        return msg;
    }

    static txTypeSwapOrder(tx, ctx){
        const msg = {
            'Type :' : [],
            'Is Buy Order :' : [],
            'Input Address :' : [],
            'Input :' : [],
            'Output Address :' : [],
            'Output :' : [],
            'Deadline :' : [],
        };
        msg['Type :'].push(tx.type);
        const nativeDenom = ctx.$store.state.nativeToken;
        if(tx && tx.msgs && Array.isArray(tx.msgs)){
            tx.msgs.forEach((m) =>{

                msg['Is Buy Order :'].push(m.msg.is_buy_order);
                msg['Input Address :'].push(m.msg.input.address);
                if(m.msg.input.coin){
                    let amount = m.msg.input.coin.amount, denom = m.msg.input.coin.denom.toUpperCase();
                    if(m.msg.input.coin.denom === nativeDenom){
                        amount = formatMsgsAndTags.formatAmount(m.msg.input.coin.amount)
                    }

                    msg['Input :'].push(`${amount} ${denom}`);
                }
                msg['Output Address :'].push(m.msg.output.address);
                if(m.msg.output.coin){
                    console.error('---',nativeDenom)
                    let amount = m.msg.output.coin.amount, denom = m.msg.output.coin.denom.toUpperCase();
                    if(m.msg.output.coin.denom === nativeDenom){
                        amount = formatMsgsAndTags.formatAmount(m.msg.output.coin.amount)
                    }

                    msg['Output :'].push(`${amount} ${denom}`);
                }
                msg['Deadline :'].push(Tools.getDisplayDate(m.msg.deadline*1000));

            })
        }
        return msg;
    }



    static txTypeTransfer(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.from_address)
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.to_address);
                    }

                    amountObj = Tools.formatAmountOfTxDetail(item.msg.amount);
                    if(amountObj && amountObj.moreAmountsNumber && amountObj.moreAmountsNumber.length > 0){
                        //handle more tokens
                        amountObj.moreAmountsNumber.forEach((item) =>{
                            message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(item.amount)} ${item.denom}`)
                        })
                    } else {
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`)
                    }

                }
            })
        }
        return message
    }

    static txTypeBurn(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.owner);
                        amountObj = Tools.formatAmountOfTxDetail(item.msg.coins);
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`);
                    }
                }
            })
        }
        return message
    }

    static txTypeSetMemoRegexp(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MEMOREGEXP] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner)
                        message[Constant.TRANSACTIONMESSAGENAME.MEMOREGEXP].unshift(item.msg.memo_regexp ? item.msg.memo_regexp : '--')
                    }
                }
            })
        }
        return message
    }

    static txTypeCreateValidator(dataTx, txType){
        let message = {}, selfBondedObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SELFBONDED] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNERADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.CONSENSUSPUBKEY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DETAILS] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        selfBondedObj = Tools.formatAmountOfTxDetail(item.msg.delegation);
                        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS].unshift(item.msg.validator_address);
                        message[Constant.TRANSACTIONMESSAGENAME.MONIKER].unshift(item.msg.valdescription.moniker ? item.msg.valdescription.moniker : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY].unshift(item.msg.valdescription.identity ? item.msg.valdescription.identity : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SELFBONDED].unshift(`${formatMsgsAndTags.formatAmount(selfBondedObj.amountNumber)} ${selfBondedObj.tokenName}`);
                        message[Constant.TRANSACTIONMESSAGENAME.OWNERADDRESS].unshift(item.msg.delegator_address);
                        message[Constant.TRANSACTIONMESSAGENAME.CONSENSUSPUBKEY].unshift(item.msg.pubkey);
                        message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE].unshift(`${Tools.formatPercentage(item.msg.commission.rate)} %`)
                        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE].unshift(item.msg.valdescription.website ? item.msg.valdescription.website : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DETAILS].unshift(item.msg.valdescription.details ? item.msg.valdescription.details : '--');
                        message.tooltip = {
                            maxChangeRate : `${Tools.formatPercentage(item.msg.commission.max_change_rate)}%`,
                            maxRate : `${Tools.formatPercentage(item.msg.commission.max_rate)} %`
                        }
                    }
                }
            })
        }
        return message
    }

    static txTypeEditValidator(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DETAILS] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS].unshift(item.msg.address);
                        message[Constant.TRANSACTIONMESSAGENAME.MONIKER].unshift(item.msg.valdescription.moniker ? item.msg.valdescription.moniker : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY].unshift(item.msg.valdescription.identity ? item.msg.valdescription.identity : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE].unshift(item.msg.valdescription.website ? item.msg.valdescription.website : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DETAILS].unshift(item.msg.valdescription.details ? item.msg.valdescription.details : '--');
                        if(dataTx.status === 'success'){
                            message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE].unshift(item.msg.commission_rate ? `${ Tools.formatPercentage(item.msg.commission_rate)} %` : '--');
                        } else {
                            message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE].unshift(item.msg.commission_rate ? `${ Tools.formatPercentage(item.msg.commission_rate)} %` : '--');
                        }
                    }
                }
            })
        }
        return message
    }

    static txTypeUnjail(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS].unshift(item.msg.address);
                    }
                }
            })
        }
        return message
    }

    static txTypeDelegate(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.delegator_addr);
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.validator_addr);
                        amountObj = Tools.formatAmountOfTxDetail(item.msg.delegation);
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`)
                    }
                }
            })
        }
        return message
    }

    static txTypeBeginRedelegate(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SHARES] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TOSHARES] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ENDTIME] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            if(dataTx.tags){
                amountObj = Tools.formatListByTagsBalance(dataTx.tags);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`);
                message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(dataTx.tags['source-validator']);
                message[Constant.TRANSACTIONMESSAGENAME.SHARES].unshift(`${Tools.numberMoveDecimal(dataTx.tags['shares-src'])} SHARES`);
                message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.tags['destination-validator'])
                message[Constant.TRANSACTIONMESSAGENAME.TOSHARES].unshift(`${Tools.numberMoveDecimal(dataTx.tags['shares-dst'])} SHARES`);
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift(Tools.format2UTC(dataTx.tags['end-time']));
            }
        } else {
            if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                dataTx.msgs.forEach(item =>{
                    if(item.type === txType){
                        if(item.msg){
                            message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--');
                            message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.validator_src_addr);
                            message[Constant.TRANSACTIONMESSAGENAME.SHARES].unshift(item.msg.shares_amount ? `${Tools.numberMoveDecimal(item.msg.shares_amount)} SHARES` : '--');
                            message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.validator_dst_addr)
                            message[Constant.TRANSACTIONMESSAGENAME.TOSHARES].unshift('--');
                            message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift('--');
                        }
                    }
                })
            }
        }
        return message
    }

    static txTypeSetWithdrawAddress(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NEWADDRESS] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALADDRESS].unshift(item.msg.delegator_addr);
                        message[Constant.TRANSACTIONMESSAGENAME.NEWADDRESS].unshift(item.msg.withdraw_addr)
                    }
                }
            })
        }
        return message
    }

    static txTypeBeginUnbonding(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SHARES] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ENDTIME] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            if(dataTx.tags){
                amountObj = Tools.formatListByTagsBalance(dataTx.tags, true);
                message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(dataTx.tags['source-validator']);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`);
                message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.tags.delegator);
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift(Tools.format2UTC(dataTx.tags['end-time']))
            }
            if(dataTx.msgs){
                if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                    dataTx.msgs.forEach(item =>{
                        if(item.msg){
                            message[Constant.TRANSACTIONMESSAGENAME.SHARES].unshift(`${Tools.numberMoveDecimal(item.msg.shares_amount)} SHARES`)
                        }
                    })
                }
            }
        } else {
            if(dataTx.tags && dataTx.tags.balance && dataTx.tags['end-time']){
                amountObj = Tools.formatListByTagsBalance(dataTx.tags);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`);
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift(Tools.format2UTC(dataTx.tags['end-time']))

            } else {
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift('--')
            }
            if(dataTx.msgs){
                if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                    dataTx.msgs.forEach(item =>{
                        if(item.msg){
                            amountObj = Tools.formatListByTagsBalance(item.msg.shares_amount, true);
                            amountObj.amountNumber !== '--' ? message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`) : message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--');
                            message[Constant.TRANSACTIONMESSAGENAME.SHARES].unshift(`${Tools.numberMoveDecimal(item.msg.shares_amount)} SHARES`)
                            message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg['validator_addr']);
                            message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.delegator_addr);
                        }
                    })
                }
            }
        }
        return message
    }

    static txTypeWithdrawDelegatorReward(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            if(dataTx.tags && dataTx.tags['withdraw-reward-total']){
                let tags = {};
                tags.balance = dataTx.tags['withdraw-reward-total']
                amountObj = Tools.formatListByTagsBalance(tags, true);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`)
            } else {
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--')
            }
        } else {
            message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--')
        }

        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.msg){
                    message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.validator_addr);
                    message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.delegator_addr)
                }
            })
        }
        return message
    }

    static txTypeWithdrawDelegatorRewardsAll(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            let formAddressArray = [], toAddressArray = [];
            if(dataTx.tags){
                for(let item in dataTx.tags){
                    if(item.startsWith('withdraw-reward-from-validator')){
                        formAddressArray.push(item.split('-')[item.split('-').length - 1])
                    }
                    if(item === 'withdraw-address'){
                        toAddressArray.push(dataTx.tags[item])
                    }
                }
            }
            message[Constant.TRANSACTIONMESSAGENAME.FROM] = formAddressArray.length > 0 ? formAddressArray : '-';
            if(dataTx.monikers){
                message[Constant.TRANSACTIONMESSAGENAME.FROM] = message[Constant.TRANSACTIONMESSAGENAME.FROM].map(item =>{
                    return {
                        address : item,
                        moniker : dataTx.monikers[item]
                    }
                })
            }
            if(dataTx.amount && dataTx.amount.length > 0){
                amountObj = Tools.formatAmountOfTxDetail(dataTx.amount);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`)
            } else {
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--')
            }
            message[Constant.TRANSACTIONMESSAGENAME.TO] = toAddressArray.length > 0 ? toAddressArray : '-';
        } else {
            if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                dataTx.msgs.forEach(item =>{
                    if(item.msg){
                        amountObj = Tools.formatAmountOfTxDetail(dataTx.amount);
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(dataTx.amount.length > 0 ? `${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}` : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.validator_addr ? item.msg.validator_addr : '-');
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.delegator_addr ? item.msg.delegator_addr : '-')
                    }
                })
            }

        }
        return message
    }

    static txTypeWithdrawValidatorRewardsAll(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        let formAddressArray = [], toAddressArray = [];
        if(dataTx.status === 'success'){
            if(dataTx.tags){
                for(let item in dataTx.tags){
                    if(item.startsWith('withdraw-reward-from-validator')){
                        formAddressArray.push(item.split('-')[item.split('-').length - 1])
                    }
                    if(item === 'withdraw-address'){
                        toAddressArray.push(dataTx.tags[item])
                    }
                }
            }
        } else {
            if(dataTx.msgs){
                dataTx.msgs.forEach(item =>{
                    if(item.msg){
                        formAddressArray.unshift(item.msg.validator_addr ? item.msg.validator_addr : '--')
                    }
                })
            }
        }
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = formAddressArray.length > 0 ? formAddressArray : '-';
        if(dataTx.monikers){
            if(message[Constant.TRANSACTIONMESSAGENAME.FROM] !== '-'){
                message[Constant.TRANSACTIONMESSAGENAME.FROM] = message[Constant.TRANSACTIONMESSAGENAME.FROM].map(item =>{
                    if(dataTx.monikers[item]){
                        return {
                            address : item,
                            moniker : dataTx.monikers[item]
                        }
                    } else {
                        return {
                            address : item,
                            moniker : dataTx.monikers[item],
                            isLink : true,
                        }
                    }

                })
            }
        }
        amountObj = Tools.formatAmountOfTxDetail(dataTx.amount);
        amountObj.amountNumber === '--' || amountObj.tokenName === '--' ? message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--') : message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`)
        message[Constant.TRANSACTIONMESSAGENAME.TO] = toAddressArray.length > 0 ? toAddressArray : '-';
        return message
    }

    static txTypeSubmitProposal(dataTx, txType){
        let message = {}, initialDepositObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TITLE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                switch (item.type){
                    case Constant.SUBMITPROPOSALTYPE.SUBMITPROPOSAL:
                        if(item.msg){
                            initialDepositObj = Tools.formatAmountOfTxDetail(item.msg.initialDeposit);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSER].unshift(item.msg.proposer);
                            message[Constant.TRANSACTIONMESSAGENAME.TITLE].unshift(item.msg.title);
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${formatMsgsAndTags.formatAmount(initialDepositObj.amountNumber)} ${initialDepositObj.tokenName}`);
                            message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.description);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE].unshift(item.msg.proposalType);
                            if(item.msg.params && item.msg.params.length > 0){
                                message[Constant.TRANSACTIONMESSAGENAME.PARAMETER] = [];
                                message[Constant.TRANSACTIONMESSAGENAME.PARAMETER].unshift(JSON.stringify(item.msg.params));
                            }
                        }
                        break;
                    case Constant.SUBMITPROPOSALTYPE.SUBMITSOFTWAREUPGRADEPROPOSAL:
                        message[Constant.TRANSACTIONMESSAGENAME.SOFTWARE] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.VERSION] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.SWITCHHEIGHT] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.TRESHOLD] = [];
                        if(item.msg){
                            initialDepositObj = Tools.formatAmountOfTxDetail(item.msg.doctxmsgsubmitproposal.initialDeposit);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSER].unshift(item.msg.doctxmsgsubmitproposal.proposer);
                            message[Constant.TRANSACTIONMESSAGENAME.TITLE].unshift(item.msg.doctxmsgsubmitproposal.title);
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${formatMsgsAndTags.formatAmount(initialDepositObj.amountNumber)} ${initialDepositObj.tokenName}`);
                            message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.doctxmsgsubmitproposal.description);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE].unshift(item.msg.doctxmsgsubmitproposal.proposalType);
                            message[Constant.TRANSACTIONMESSAGENAME.SOFTWARE].unshift(item.msg.software);
                            message[Constant.TRANSACTIONMESSAGENAME.VERSION].unshift(item.msg.version);
                            message[Constant.TRANSACTIONMESSAGENAME.SWITCHHEIGHT].unshift(item.msg.switch_height);
                            message[Constant.TRANSACTIONMESSAGENAME.TRESHOLD].unshift(`${Tools.formatPercentage(item.msg.threshold)} %`);
                        }
                        break;
                    case Constant.SUBMITPROPOSALTYPE.SUBMITTOKENADDITIONPROPOSAL:
                        message[Constant.TRANSACTIONMESSAGENAME.SYMBOL] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.NAME] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.DECIMAL] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.MINUNITALIAS] = [];
                        if(item.msg){
                            initialDepositObj = Tools.formatAmountOfTxDetail(item.msg.doctxmsgsubmitproposal.initialDeposit);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSER].unshift(item.msg.doctxmsgsubmitproposal.proposer);
                            message[Constant.TRANSACTIONMESSAGENAME.TITLE].unshift(item.msg.doctxmsgsubmitproposal.title);
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${formatMsgsAndTags.formatAmount(initialDepositObj.amountNumber)} ${initialDepositObj.tokenName}`);
                            message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.doctxmsgsubmitproposal.description);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE].unshift(item.msg.doctxmsgsubmitproposal.proposalType);

                            message[Constant.TRANSACTIONMESSAGENAME.SYMBOL] = [item.msg.symbol];
                            message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL] = [item.msg.canonical_symbol];
                            message[Constant.TRANSACTIONMESSAGENAME.NAME] = [item.msg.name];
                            message[Constant.TRANSACTIONMESSAGENAME.DECIMAL] = [item.msg.decimal];
                            message[Constant.TRANSACTIONMESSAGENAME.MINUNITALIAS] = [item.msg.min_unit_alias];
                        }
                        break;
                    case Constant.SUBMITPROPOSALTYPE.SUBMITTXTAXUSAGEPROPOSAL:
                        message[Constant.TRANSACTIONMESSAGENAME.USAGE] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.DESTADDRESS] = [];
                        message[Constant.TRANSACTIONMESSAGENAME.PERCENT] = [];
                        if(item.msg){
                            initialDepositObj = Tools.formatAmountOfTxDetail(item.msg.doctxmsgsubmitproposal.initialDeposit);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSER].unshift(item.msg.doctxmsgsubmitproposal.proposer);
                            message[Constant.TRANSACTIONMESSAGENAME.TITLE].unshift(item.msg.doctxmsgsubmitproposal.title);
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${formatMsgsAndTags.formatAmount(initialDepositObj.amountNumber)} ${initialDepositObj.tokenName}`);
                            message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.doctxmsgsubmitproposal.description);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE].unshift(item.msg.doctxmsgsubmitproposal.proposalType);
                            message[Constant.TRANSACTIONMESSAGENAME.USAGE].unshift(item.msg.usage);
                            message[Constant.TRANSACTIONMESSAGENAME.PERCENT].unshift(`${numberMoveDecimal(item.msg.percent, 2)} %`);
                            if(item.msg.usage === Constant.TxType.BURN){
                                message[Constant.TRANSACTIONMESSAGENAME.DESTADDRESS].unshift('-');
                            } else {
                                message[Constant.TRANSACTIONMESSAGENAME.DESTADDRESS].unshift(item.msg.dest_address);
                            }
                        }

                }
            })
        }
        return message
    }

    static txTypeDeposit(dataTx, txType){
        let message = {}, depositAmoutObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DEPOSITOR] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DEPOSIT] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        depositAmoutObj = Tools.formatAmountOfTxDetail(item.msg.amount)
                        message[Constant.TRANSACTIONMESSAGENAME.DEPOSITOR].unshift(item.msg.depositor)
                        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALID].unshift(item.msg.proposal_id)
                        message[Constant.TRANSACTIONMESSAGENAME.DEPOSIT].unshift(`${formatMsgsAndTags.formatAmount(depositAmoutObj.amountNumber)} ${depositAmoutObj.tokenName}`)
                    }
                }
            })
        }
        return message
    }

    static txTypeVote(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.VOTER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OPTION] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.VOTER].unshift(item.msg.voter)
                        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALID].unshift(item.msg.proposal_id)
                        message[Constant.TRANSACTIONMESSAGENAME.OPTION].unshift(item.msg.option)
                    }
                }
            })
        }
        return message
    }

    static txTypeIssueToken(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        //message[Constant.TRANSACTIONMESSAGENAME.FAMILY] = [];
        //message[Constant.TRANSACTIONMESSAGENAME.SOURCE] = [];
        //message[Constant.TRANSACTIONMESSAGENAME.GATEWAY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SYMBOL] = [];
        //message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NAME] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DECIMAL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SYMBOLMINALIAS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.INITIALSUPPLY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        //message[Constant.TRANSACTIONMESSAGENAME.FAMILY].unshift(item.msg.family ? item.msg.family : '--');
                        //message[Constant.TRANSACTIONMESSAGENAME.SOURCE].unshift(item.msg.source ? item.msg.source : '--');
                        //message[Constant.TRANSACTIONMESSAGENAME.GATEWAY].unshift(item.msg.gateway ? item.msg.gateway : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SYMBOL].unshift(item.msg.symbol ? item.msg.symbol : '--');
                        //message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL].unshift(item.msg.canonical_symbol ? item.msg.canonical_symbol : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.NAME].unshift(item.msg.name ? item.msg.name : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DECIMAL].unshift((item.msg.scale || item.msg.scale === 0) ? String(item.msg.scale) : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SYMBOLMINALIAS].unshift(item.msg.min_unit ? item.msg.min_unit : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.INITIALSUPPLY].unshift(item.msg.initial_supply || item.msg.initial_supply === 0 ? formatMsgsAndTags.formatAmount(item.msg.initial_supply) : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY].unshift(item.msg.max_supply || item.msg.max_supply === 0 ? formatMsgsAndTags.formatAmount(item.msg.max_supply) : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE].unshift(item.msg.mintable);
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner ? item.msg.owner : '--');
                    }
                }
            })
        }
        return message
    }

    static txTypeEditToken(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NAME] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL].unshift(item.msg.symbol ? item.msg.symbol : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.NAME].unshift(item.msg.name ? item.msg.name : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY].unshift(item.msg.max_supply || item.msg.max_supply === 0 ? formatMsgsAndTags.formatAmount(item.msg.max_supply) : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE].unshift(item.msg.mintable);
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner ? item.msg.owner : '--');
                    }
                }
            })
        }
        return message
    }

    static txTypeMintToken(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner ? item.msg.owner : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(formatMsgsAndTags.formatAmount(item.msg.amount));
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.to);
                    }
                }
            })
        }
        return message
    }

    static txTypeTransferTokenOwner(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALOWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NEWOWNER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALOWNER].unshift(item.msg.src_owner);
                        message[Constant.TRANSACTIONMESSAGENAME.NEWOWNER].unshift(item.msg.dst_owner);
                    }
                }
            })
        }
        return message
    }

    static txTypeCreateGateway(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DETAILS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner);
                        message[Constant.TRANSACTIONMESSAGENAME.MONIKER].unshift(item.msg.moniker ? item.msg.moniker : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY].unshift(item.msg.identity ? item.msg.identity : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DETAILS].unshift(item.msg.details ? item.msg.details : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE].unshift(item.msg.website ? item.msg.website : '--');
                    }
                }
            })
        }
        return message

    }

    static txTypeEditGateway(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DETAILS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner);
                        message[Constant.TRANSACTIONMESSAGENAME.MONIKER].unshift(item.msg.moniker ? item.msg.moniker : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY].unshift(item.msg.identity ? item.msg.identity : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DETAILS].unshift(item.msg.details ? item.msg.details : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE].unshift(item.msg.website ? item.msg.website : '--');
                    }
                }
            })
        }
        return message
    }

    static txTypeTransferGatewayOwner(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALOWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NEWOWNER] = [];

        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALOWNER].unshift(item.msg.owner ? item.msg.owner : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MONIKER].unshift(item.msg.moniker ? item.msg.moniker : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.NEWOWNER].unshift(item.msg.to ? item.msg.to : '--');
                    }
                }
            })
        }
        return message
    }

    static txTypeRequestRand(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.BLOCKINTERVAL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.REQUESTID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.RANDHEIGHT] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.BLOCKINTERVAL].unshift(item.msg['block-interval'] || item.msg['block-interval'] === 0 ? item.msg['block-interval'] : '--');
                    }
                }
            })
        }
        if(dataTx.tags && JSON.stringify(Object.keys(dataTx.tags)[0]) !== '{}'){
            message[Constant.TRANSACTIONMESSAGENAME.RANDHEIGHT].unshift(dataTx.tags['rand-height'] || dataTx.tags['rand-height'] == 0 ? dataTx.tags['rand-height'] : '--')
            message[Constant.TRANSACTIONMESSAGENAME.REQUESTID].unshift(dataTx.tags['request-id'] ? dataTx.tags['request-id'] : '--')
        }
        return message
    }

    static txTypeAddProfiler(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESSBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.addguardian.description ? item.msg.addguardian.description : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS].unshift(item.msg.addguardian.address ? item.msg.addguardian.address : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.ADDRESSBY].unshift(item.msg.addguardian.added_by ? item.msg.addguardian.added_by : '--');
                    }
                }
            })
        }
        return message
    }

    static txTypeAddTrustee(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESSBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.addguardian.description ? item.msg.addguardian.description : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS].unshift(item.msg.addguardian.address ? item.msg.addguardian.address : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.ADDRESSBY].unshift(item.msg.addguardian.added_by ? item.msg.addguardian.added_by : '--');
                    }
                }
            })
        }
        return message
    }

    static txTypeDeleteProfiler(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DELETEDBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS].unshift(item.msg.deleteguardian.address ? item.msg.deleteguardian.address : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DELETEDBY].unshift(item.msg.deleteguardian.deleted_by ? item.msg.deleteguardian.deleted_by : '--');
                    }
                }
            })
        }
        return message

    }

    static txTypeDeleteTrustee(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DELETEDBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS].unshift(item.msg.deleteguardian.address ? item.msg.deleteguardian.address : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DELETEDBY].unshift(item.msg.deleteguardian.deleted_by ? item.msg.deleteguardian.deleted_by : '--');
                    }
                }
            })
        }
        return message

    }

    static txTypeClaimHTLC(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SECRET] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            if(dataTx.tags){
                message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(dataTx.tags.sender);
                message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.tags.receiver);
                message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(dataTx.tags['hash-lock']);
                message[Constant.TRANSACTIONMESSAGENAME.SECRET].unshift(dataTx.tags.secret);
            }
        } else {
            if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                dataTx.msgs.forEach(item =>{
                    if(item.type === txType){
                        if(item.msg){
                            message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(item.msg['hash_lock']);
                            message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift('-');
                            message[Constant.TRANSACTIONMESSAGENAME.TO].unshift('-');
                            message[Constant.TRANSACTIONMESSAGENAME.SECRET].unshift(item.msg.secret);
                        }
                    }
                })
            }
        }

        return message
    }

    static txTypeCreateHTLC(dataTx, txType){
        let message = {}, amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TIMELOCK] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TIMESTAMP] = [];
        message[Constant.TRANSACTIONMESSAGENAME.EXPIRYHEIGHT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.CROSSCHAINREVEIVER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item =>{
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        amountObj = Tools.formatAmountOfTxDetail(item.msg.amount);
                        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(item.msg.hash_lock);
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.sender);
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${formatMsgsAndTags.formatAmount(amountObj.amountNumber)} ${amountObj.tokenName}`);
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.to);
                        message[Constant.TRANSACTIONMESSAGENAME.TIMELOCK].unshift(item.msg.time_lock || item.msg.time_lock == 0 ? item.msg.time_lock : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.TIMESTAMP].unshift(item.msg.timestamp);
                        message[Constant.TRANSACTIONMESSAGENAME.EXPIRYHEIGHT].unshift(dataTx.expire_height > 0 ? dataTx.expire_height : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.CROSSCHAINREVEIVER].unshift(item.msg['receiver_on_other_chain'] ? item.msg['receiver_on_other_chain'] : '--');
                    }
                }
            })
        }
        return message
    }

    static txTypeRefundHTLC(dataTx, txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            if(dataTx.status === 'success'){
                message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
                if(dataTx.tags){
                    message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(dataTx.tags['hash-lock']);
                    message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.tags.sender);
                }
            } else {
                if(dataTx.msgs){
                    dataTx.msgs.forEach(item =>{
                        if(item.type === txType){
                            message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                            if(item.msg){
                                message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(item.msg['hash_lock']);
                            }
                        }
                    })
                }
                if(dataTx.tags && dataTx.tags.sender){
                    message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.tags.sender ? dataTx.tags.sender : '-');
                } else {
                    message[Constant.TRANSACTIONMESSAGENAME.TO].unshift('-');
                }
            }

        }
        return message
    }
}
