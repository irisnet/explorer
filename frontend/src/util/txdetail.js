
import Constant from "../constant/Constant"
import numberMoveDecimal from "move-decimal-point"
import Tools from "../util/Tools"
export default class formatMsgsAndTags {
    static switchTxType(dataTx){
        let message;
        switch (dataTx.type) {
            case Constant.TxType.TRANSFER :
                return message = formatMsgsAndTags.txTypeTransfer(dataTx,Constant.TxType.TRANSFER);
                break;
            case Constant.TxType.BURN :
                return message = formatMsgsAndTags.txTypeBurn(dataTx,Constant.TxType.BURN);
                break;
            case Constant.TxType.SETMEMOREGEXP :
                return message = formatMsgsAndTags.txTypeSetMemoRegexp(dataTx,Constant.TxType.SETMEMOREGEXP);
                break;
            case Constant.TxType.CREATEVALIDATOR :
                return message = formatMsgsAndTags.txTypeCreateValidator(dataTx,Constant.TxType.CREATEVALIDATOR);
                break;
            case Constant.TxType.EDITVALIDATOR :
                return message = formatMsgsAndTags.txTypeEditValidator(dataTx,Constant.TxType.EDITVALIDATOR);
                break;
            case Constant.TxType.UNJAIL :
                return message = formatMsgsAndTags.txTypeUnjail(dataTx,Constant.TxType.UNJAIL);
                break;
            case Constant.TxType.DELEGATE :
                return message = formatMsgsAndTags.txTypeDelegate(dataTx,Constant.TxType.DELEGATE);
                break;
            case Constant.TxType.BEGINREDELEGATE :
                return message = formatMsgsAndTags.txTypeBeginRedelegate(dataTx,Constant.TxType.BEGINREDELEGATE);
                break;
            case Constant.TxType.SETWITHDRAWADDRESS :
                return message = formatMsgsAndTags.txTypeSetWithdrawAddress(dataTx,Constant.TxType.SETWITHDRAWADDRESS);
                break;
            case Constant.TxType.BEGINUNBONDING :
                return message = formatMsgsAndTags.txTypeBeginUnbonding(dataTx,Constant.TxType.BEGINUNBONDING);
                break;
            case Constant.TxType.WITHDRAWDELEGATORREWARD :
                return message = formatMsgsAndTags.txTypeWithdrawDelegatorReward(dataTx,Constant.TxType.WITHDRAWDELEGATORREWARD);
                break;
            case Constant.TxType.WITHDRAWDELEGATORREWARDSALL :
                return message = formatMsgsAndTags.txTypeWithdrawDelegatorRewardsAll(dataTx,Constant.TxType.WITHDRAWDELEGATORREWARDSALL);
                break;
            case Constant.TxType.WITHDRAWVALIDATORREWARDSALL :
                return message = formatMsgsAndTags.txTypeWithdrawValidatorRewardsAll(dataTx,Constant.TxType.WITHDRAWVALIDATORREWARDSALL);
                break;
            case Constant.TxType.SUBMITPROPOSAL :
                return message = formatMsgsAndTags.txTypeSubmitProposal(dataTx,Constant.TxType.SUBMITPROPOSAL);
                break;
            case Constant.TxType.DEPOSIT :
                return message = formatMsgsAndTags.txTypeDeposit(dataTx,Constant.TxType.DEPOSIT);
                break;
            case Constant.TxType.VOTE :
                return message = formatMsgsAndTags.txTypeVote(dataTx,Constant.TxType.VOTE);
                break;
            case Constant.TxType.ISSUETOKEN :
                return message = formatMsgsAndTags.txTypeIssueToken(dataTx,Constant.TxType.ISSUETOKEN);
                break;
            case Constant.TxType.EDITTOKEN :
                return message = formatMsgsAndTags.txTypeEditToken(dataTx,Constant.TxType.EDITTOKEN);
                break;
            case Constant.TxType.MINTTOKEN :
                return message = formatMsgsAndTags.txTypeMintToken(dataTx,Constant.TxType.MINTTOKEN);
                break;
            case Constant.TxType.TRANSFERTOKENOWNER :
                return message = formatMsgsAndTags.txTypeTransferTokenOwner(dataTx,Constant.TxType.TRANSFERTOKENOWNER);
                break;
            case Constant.TxType.CREATEGATEWAY :
                return message = formatMsgsAndTags.txTypeCreateGateway(dataTx,Constant.TxType.CREATEGATEWAY);
                break;
            case Constant.TxType.EDITGATEWAY :
                return message = formatMsgsAndTags.txTypeEditGateway(dataTx,Constant.TxType.EDITGATEWAY);
                break;
            case Constant.TxType.TRANSFERGATEWAYOWNER :
                return message = formatMsgsAndTags.txTypeTransferGatewayOwner(dataTx,Constant.TxType.TRANSFERGATEWAYOWNER);
                break;
            case Constant.TxType.REQUESTRAND :
                return message = formatMsgsAndTags.txTypeRequestRand(dataTx,Constant.TxType.REQUESTRAND);
                break;
            case Constant.TxType.ADDPROFILER :
                return message = formatMsgsAndTags.txTypeAddProfiler(dataTx,Constant.TxType.ADDPROFILER);
                break;
            case Constant.TxType.ADDTRUSTEE :
                return message = formatMsgsAndTags.txTypeAddTrustee(dataTx,Constant.TxType.ADDTRUSTEE);
                break;
            case Constant.TxType.DELETEPROFIKER :
                return message = formatMsgsAndTags.txTypeDeleteProfiler(dataTx,Constant.TxType.DELETEPROFIKER);
                break;
            case Constant.TxType.DELETETRUSTEE :
                return message = formatMsgsAndTags.txTypeDeleteTrustee(dataTx,Constant.TxType.DELETETRUSTEE);
                break;
            case Constant.TxType.CLAIMHTLC :
                return message = formatMsgsAndTags.txTypeClaimHTLC(dataTx,Constant.TxType.CLAIMHTLC);
                break;
            case Constant.TxType.CREATEHTLC :
                return message = formatMsgsAndTags.txTypeCreateHTLC(dataTx,Constant.TxType.CREATEHTLC);
                break;
            case Constant.TxType.REFUNDHTLC :
                return message = formatMsgsAndTags.txTypeRefundHTLC(dataTx,Constant.TxType.REFUNDHTLC);
                break;
            case Constant.TxType.ADDLIQUIDITY :
                return message = formatMsgsAndTags.txTypeAddLiquidity(dataTx,Constant.TxType.ADDLIQUIDITY);
                break;
            case Constant.TxType.REMOVELIQUIDITY :
                return message = formatMsgsAndTags.txTypeRemoveLiquidity(dataTx,Constant.TxType.REMOVELIQUIDITY);
                break;
            case Constant.TxType.SWAPORDER :
                return message = formatMsgsAndTags.txTypeSwapOrder(dataTx,Constant.TxType.SWAPORDER);
        }

    }
    static txTypeTransfer(dataTx,txType){
        let message = {},amountObj;
            message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
            message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
            message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
            message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg && item.msg.inputs && Array.isArray(item.msg.inputs)){
                        item.msg.inputs.forEach(item => {
                            message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.address)
                        })
                    }
                    if(item.msg && item.msg.outputs && Array.isArray(item.msg.outputs)){
                        item.msg.outputs.forEach(item => {
                            message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.address);
                            amountObj = Tools.formatAmountOfTxDetail(item.coins);
                            if(amountObj && amountObj.moreAmountsNumber && amountObj.moreAmountsNumber.length > 0){
                                //handle more tokens
                                amountObj.moreAmountsNumber.forEach( (item) => {
                                    message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${item.amount} ${item.denom}`)
                                })
                            }else {
                                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`)
                            }
                        });
                    }
                }
            })
        }
        return message
    }
    static txTypeBurn(dataTx,txType){
        let message = {},amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.owner);
                        amountObj = Tools.formatAmountOfTxDetail(item.msg.coins);
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`);
                    }
                }
            })
        }
        return message
    }
    static txTypeSetMemoRegexp(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MEMOREGEXP] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeCreateValidator(dataTx,txType){
        let message = {},selfBondedObj;
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
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        selfBondedObj = Tools.formatAmountOfTxDetail(item.msg.delegation);
                        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS].unshift(item.msg.validator_address);
                        message[Constant.TRANSACTIONMESSAGENAME.MONIKER].unshift(item.msg.valdescription.moniker ? item.msg.valdescription.moniker : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY].unshift(item.msg.valdescription.identity ? item.msg.valdescription.identity : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SELFBONDED].unshift(`${selfBondedObj.amountNumber} ${selfBondedObj.tokenName}`);
                        message[Constant.TRANSACTIONMESSAGENAME.OWNERADDRESS].unshift(item.msg.delegator_address);
                        message[Constant.TRANSACTIONMESSAGENAME.CONSENSUSPUBKEY].unshift(item.msg.pubkey);
                        message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE].unshift(`${Number(item.msg.commission.rate * 100)} %`)
                        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE].unshift(item.msg.valdescription.website ? item.msg.valdescription.website : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DETAILS].unshift(item.msg.valdescription.details ? item.msg.valdescription.details : '--');
                        message.tooltip = {
                            maxChangeRate:  `${Number(item.msg.commission.max_change_rate * 100)}%`,
                            maxRate: `${Number(item.msg.commission.max_rate) * 100} %`
                        }
                    }
                }
            })
        }
        return message
    }
    static txTypeEditValidator(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DETAILS] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS].unshift(item.msg.address);
                        message[Constant.TRANSACTIONMESSAGENAME.MONIKER].unshift(item.msg.valdescription.moniker ? item.msg.valdescription.moniker : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY].unshift(item.msg.valdescription.identity ? item.msg.valdescription.identity : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE].unshift(item.msg.valdescription.website ? item.msg.valdescription.website : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DETAILS].unshift(item.msg.valdescription.details ? item.msg.valdescription.details : '--');
                        if(dataTx.status === 'success'){
                            message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE].unshift(item.msg.commission ? `${ Number(item.msg.commission.rate * 100)} %` : '--');
                        }else {
                            message[Constant.TRANSACTIONMESSAGENAME.COMMISSIONRATE].unshift(item.msg.commission_rate ? `${ Number(item.msg.commission_rate * 100)} %` : '--');
                        }
                    }
                }
            })
        }
        return message
    }
    static txTypeUnjail(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OPERATORADDRESS] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeDelegate(dataTx,txType){
        let message = {},amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.delegator_addr);
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.validator_addr);
                        amountObj = Tools.formatAmountOfTxDetail(item.msg.delegation);
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`)
                    }
                }
            })
        }
        return message
    }
    static txTypeBeginRedelegate(dataTx,txType){
        let message = {},amountObj;
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
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`);
                message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(dataTx.tags['source-validator']);
                message[Constant.TRANSACTIONMESSAGENAME.SHARES].unshift(`${Tools.numberMoveDecimal(dataTx.tags['shares-src'])} SHARES`);
                message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.tags['destination-validator'])
                message[Constant.TRANSACTIONMESSAGENAME.TOSHARES].unshift(`${Tools.numberMoveDecimal(dataTx.tags['shares-dst'])} SHARES`);
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift(Tools.format2UTC(dataTx.tags['end-time']));
            }
        }else {
            if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                dataTx.msgs.forEach( item => {
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
    static txTypeSetWithdrawAddress(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NEWADDRESS] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeBeginUnbonding(dataTx,txType){
        let message = {},amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ENDTIME] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            if(dataTx.tags){
                amountObj = Tools.formatListByTagsBalance(dataTx.tags,true);
                message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(dataTx.tags['source-validator']);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`);
                message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.tags.delegator);
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift(Tools.format2UTC(dataTx.tags['end-time']))
            }
        }else {
            if(dataTx.tags && dataTx.tags.balance && dataTx.tags['end-time']){
                amountObj = Tools.formatListByTagsBalance(dataTx.tags);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`);
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift(Tools.format2UTC(dataTx.tags['end-time']))

            }else {
                message[Constant.TRANSACTIONMESSAGENAME.ENDTIME].unshift('--')
            }
            if(dataTx.msgs){
                if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                    dataTx.msgs.forEach(item => {
                        if(item.msg){
                            amountObj = Tools.formatListByTagsBalance(item.msg.shares_amount,true);
                            amountObj.amountNumber !== '--' ? message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`) : message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--');
                            message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg['validator_addr']);
                            message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.delegator_addr);
                        }
                    })
                }
            }
        }
        return message
    }
    static txTypeWithdrawDelegatorReward(dataTx,txType){
        let message = {},amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            if(dataTx.tags && dataTx.tags['withdraw-reward-total']){
                let tags = {};
                tags.balance = dataTx.tags['withdraw-reward-total']
                amountObj = Tools.formatListByTagsBalance(tags,true);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`)
            }else {
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--')
            }
        }else {
            message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--')
        }

        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach(item => {
                if(item.msg){
                    message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.validator_addr);
                    message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.delegator_addr)
                }
            })
        }
        return message
    }
    static txTypeWithdrawDelegatorRewardsAll(dataTx,txType){
        let message = {},amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.status === 'success'){
            message[Constant.TRANSACTIONMESSAGENAME.FROM] = dataTx.from ? dataTx.from.split(',') : '-';
            if(dataTx.amount && dataTx.amount.length > 0 ){
                amountObj = Tools.formatAmountOfTxDetail(dataTx.amount);
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`)
            }else {
                message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--')
            }
            message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(dataTx.to);
        }else {
            if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
                dataTx.msgs.forEach(item => {
                    if(item.msg){
                        amountObj = Tools.formatAmountOfTxDetail(dataTx.amount) ;
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(dataTx.amount.length > 0 ? `${amountObj.amountNumber} ${amountObj.tokenName}` : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.validator_addr ? item.msg.validator_addr : '-');
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.delegator_addr ? item.msg.delegator_addr : '-')
                    }
                })
            }

        }
        return message
    }
    static txTypeWithdrawValidatorRewardsAll(dataTx,txType){
        let message = {},amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = dataTx.from ? dataTx.from.split(',') : '-';
        amountObj = Tools.formatAmountOfTxDetail(dataTx.amount);
        amountObj.amountNumber === '--' || amountObj.tokenName === '--' ? message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift('--') : message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`)
        message[Constant.TRANSACTIONMESSAGENAME.TO] = (dataTx.to ? dataTx.to.split(',') : '-');
        return message
    }
    static txTypeSubmitProposal(dataTx,txType){
        let message = {},initialDepositObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TITLE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(txType);
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                switch (item.type) {
                    case Constant.SUBMITPROPOSALTYPE.SUBMITPROPOSAL:
                        if(item.msg){
                            initialDepositObj = Tools.formatAmountOfTxDetail(item.msg.initialDeposit);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSER].unshift(item.msg.proposer);
                            message[Constant.TRANSACTIONMESSAGENAME.TITLE].unshift(item.msg.title);
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${initialDepositObj.amountNumber} ${initialDepositObj.tokenName}`);
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
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${initialDepositObj.amountNumber} ${initialDepositObj.tokenName}`);
                            message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.doctxmsgsubmitproposal.description);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE].unshift(item.msg.doctxmsgsubmitproposal.proposalType);
                            message[Constant.TRANSACTIONMESSAGENAME.SOFTWARE].unshift(item.msg.software);
                            message[Constant.TRANSACTIONMESSAGENAME.VERSION].unshift(item.msg.version);
                            message[Constant.TRANSACTIONMESSAGENAME.SWITCHHEIGHT].unshift(item.msg.switch_height);
                            message[Constant.TRANSACTIONMESSAGENAME.TRESHOLD].unshift(`${Number(item.msg.threshold * 100)} %`);
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
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${initialDepositObj.amountNumber} ${initialDepositObj.tokenName}`);
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
                            message[Constant.TRANSACTIONMESSAGENAME.INITIALDEPOSIT].unshift(`${initialDepositObj.amountNumber} ${initialDepositObj.tokenName}`);
                            message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION].unshift(item.msg.doctxmsgsubmitproposal.description);
                            message[Constant.TRANSACTIONMESSAGENAME.PROPOSALTYPE].unshift(item.msg.doctxmsgsubmitproposal.proposalType);
                            message[Constant.TRANSACTIONMESSAGENAME.USAGE].unshift(item.msg.usage);
                            message[Constant.TRANSACTIONMESSAGENAME.PERCENT].unshift(`${numberMoveDecimal(item.msg.percent,2)} %`);
                            if(item.msg.usage === Constant.TxType.BURN){
                                message[Constant.TRANSACTIONMESSAGENAME.DESTADDRESS].unshift('-');
                            }else {
                                message[Constant.TRANSACTIONMESSAGENAME.DESTADDRESS].unshift(item.msg.dest_address);
                            }
                        }

                }
            })
        }
        return message
    }
    static txTypeDeposit(dataTx,txType){
        let message = {},depositAmoutObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DEPOSITOR] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DEPOSIT] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        depositAmoutObj = Tools.formatAmountOfTxDetail(item.msg.amount)
                        message[Constant.TRANSACTIONMESSAGENAME.DEPOSITOR].unshift(item.msg.depositor)
                        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALID].unshift(item.msg.proposal_id)
                        message[Constant.TRANSACTIONMESSAGENAME.DEPOSIT].unshift(`${depositAmoutObj.amountNumber} ${depositAmoutObj.tokenName}`)
                    }
                }
            })
        }
        return message
    }
    static txTypeVote(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.VOTER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.PROPOSALID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OPTION] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeIssueToken(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FAMILY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SOURCE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.GATEWAY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SYMBOL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NAME] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DECIMAL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SYMBOLMINALIAS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.INITIALSUPPLY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.FAMILY].unshift(item.msg.family ? item.msg.family : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SOURCE].unshift(item.msg.source ? item.msg.source : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.GATEWAY].unshift(item.msg.gateway ? item.msg.gateway : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SYMBOL].unshift(item.msg.symbol ? item.msg.symbol : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL].unshift(item.msg.canonical_symbol ? item.msg.canonical_symbol : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.NAME].unshift(item.msg.name ? item.msg.name : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.DECIMAL].unshift(item.msg.decimal  || item.msg.decimal === 0 ? item.msg.decimal : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SYMBOLMINALIAS].unshift(item.msg.min_unit_alias? item.msg.min_unit_alias : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.INITIALSUPPLY].unshift(item.msg.initial_supply || item.msg.initial_supply === 0 ? item.msg.initial_supply : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY].unshift(item.msg.max_supply || item.msg.max_supply === 0 ? item.msg.max_supply : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE].unshift(item.msg.mintable);
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner ? item.msg.owner : '--');
                    }
                }
            })
        }
        return message
    }
    static txTypeEditToken(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TOKENID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NAME] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SYMBOLMINALIAS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.TOKENID].unshift(item.msg.token_id );
                        message[Constant.TRANSACTIONMESSAGENAME.CANONICALSYMBOL].unshift(item.msg.canonical_symbol ? item.msg.canonical_symbol : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.NAME].unshift(item.msg.name ? item.msg.name : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.SYMBOLMINALIAS].unshift(item.msg.min_unit_alias? item.msg.min_unit_alias : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MAXSUPPLY].unshift(item.msg.max_supply || item.msg.max_supply === 0 ? item.msg.max_supply : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.MINTABLE].unshift(item.msg.mintable);
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner ? item.msg.owner : '--');
                    }
                }
            })
        }
        return message
    }
    static txTypeMintToken(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TOKENID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TO] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.TOKENID].unshift(item.msg.token_id );
                        message[Constant.TRANSACTIONMESSAGENAME.OWNER].unshift(item.msg.owner ? item.msg.owner : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(item.msg.amount);
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.to);
                    }
                }
            })
        }
        return message
    }
    static txTypeTransferTokenOwner(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.TOKENID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALOWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NEWOWNER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.TOKENID].unshift(item.msg.token_id );
                        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALOWNER].unshift(item.msg.src_owner);
                        message[Constant.TRANSACTIONMESSAGENAME.NEWOWNER].unshift(item.msg.dst_owner);
                    }
                }
            })
        }
        return message
    }
    static txTypeCreateGateway(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DETAILS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeEditGateway(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.IDENTITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DETAILS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WEBSITE] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeTransferGatewayOwner(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ORIGINALOWNER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MONIKER] = [];
        message[Constant.TRANSACTIONMESSAGENAME.NEWOWNER] = [];

        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeRequestRand(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.BLOCKINTERVAL] = [];
        message[Constant.TRANSACTIONMESSAGENAME.REQUESTID] = [];
        message[Constant.TRANSACTIONMESSAGENAME.RANDHEIGHT] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.BLOCKINTERVAL].unshift(item.msg['block-interval']  || item.msg['block-interval'] === 0 ? item.msg['block-interval'] : '--');
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
    static txTypeAddProfiler(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESSBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeAddTrustee(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DESCRIPTION] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESSBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeDeleteProfiler(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DELETEDBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeDeleteTrustee(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DELETEDBY] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
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
    static txTypeClaimHTLC(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SECRET] = [];

        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(item.msg['hash_lock']);
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.sender);
                        message[Constant.TRANSACTIONMESSAGENAME.SECRET].unshift(item.msg.secret);
                    }
                }
            })
        }
        return message
    }
    static txTypeCreateHTLC(dataTx,txType){
        let message = {},amountObj;
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
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        amountObj = Tools.formatAmountOfTxDetail(item.msg.amount);
                        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(item.msg.hash_lock);
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.sender);
                        message[Constant.TRANSACTIONMESSAGENAME.AMOUNT].unshift(`${amountObj.amountNumber} ${amountObj.tokenName}`);
                        message[Constant.TRANSACTIONMESSAGENAME.TO].unshift(item.msg.to);
                        message[Constant.TRANSACTIONMESSAGENAME.TIMELOCK].unshift(item.msg.time_lock || item.msg.time_lock == 0 ? item.msg.time_lock : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.TIMESTAMP].unshift(item.msg.timestamp);
                        message[Constant.TRANSACTIONMESSAGENAME.EXPIRYHEIGHT].unshift(dataTx.expire_height > 0 ? dataTx.expire_height : '--');
                        message[Constant.TRANSACTIONMESSAGENAME.CROSSCHAINREVEIVER].unshift(item.msg['receiver_on_other_chain']);
                    }
                }
            })
        }
        return message
    }
    static txTypeRefundHTLC(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.FROM] = [];
        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.HASHLOCK].unshift(item.msg['hash_lock']);
                        message[Constant.TRANSACTIONMESSAGENAME.FROM].unshift(item.msg.sender);
                    }
                }
            })
        }
        return message
    }
    static txTypeAddLiquidity(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MAXTOKEN] = [];
        message[Constant.TRANSACTIONMESSAGENAME.EXACTIRISAMT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MINLIQUIDITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DEADLIN] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SENDER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.MAXTOKEN].unshift(`${item.msg.max_token.amount} ${item.msg.max_token.denom.toUpperCase()}`);
                        message[Constant.TRANSACTIONMESSAGENAME.EXACTIRISAMT].unshift(item.msg.exact_iris_amt);
                        message[Constant.TRANSACTIONMESSAGENAME.MINLIQUIDITY].unshift(item.msg.min_liquidity);
                        message[Constant.TRANSACTIONMESSAGENAME.DEADLIN].unshift(item.msg.deadline);
                        message[Constant.TRANSACTIONMESSAGENAME.SENDER].unshift(item.msg.sender);
                    }
                }
            })
        }
        return message
    }
    static txTypeRemoveLiquidity(dataTx,txType){
        let message = {};
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MINTOKEN] = [];
        message[Constant.TRANSACTIONMESSAGENAME.WITHDRAWLIQUIDITY] = [];
        message[Constant.TRANSACTIONMESSAGENAME.MINIRISAMT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DEADLINE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.SENDER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.MINTOKEN].unshift(item.msg.min_token);
                        message[Constant.TRANSACTIONMESSAGENAME.WITHDRAWLIQUIDITY].unshift(`${item.msg.withdraw_liquidity.amount} ${item.msg.withdraw_liquidity.denom.toUpperCase()}`);
                        message[Constant.TRANSACTIONMESSAGENAME.MINIRISAMT].unshift(item.msg.min_iris_amt);
                        message[Constant.TRANSACTIONMESSAGENAME.DEADLINE].unshift(item.msg.deadline);
                        message[Constant.TRANSACTIONMESSAGENAME.SENDER].unshift(item.msg.sender);
                    }
                }
            })
        }
        return message
    }
    static txTypeSwapOrder(dataTx,txType){
        let message = {},amountObj;
        message[Constant.TRANSACTIONMESSAGENAME.TXTYPE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.INPUTADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OUTPUTADDRESS] = [];
        message[Constant.TRANSACTIONMESSAGENAME.INPUT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.OUTPUT] = [];
        message[Constant.TRANSACTIONMESSAGENAME.DEADLINE] = [];
        message[Constant.TRANSACTIONMESSAGENAME.ISBUYORDER] = [];
        if(dataTx.msgs && Array.isArray(dataTx.msgs) && dataTx.msgs !== null){
            dataTx.msgs.forEach( item => {
                if(item.type === txType){
                    message[Constant.TRANSACTIONMESSAGENAME.TXTYPE].unshift(item.type);
                    if(item.msg){
                        message[Constant.TRANSACTIONMESSAGENAME.INPUTADDRESS].unshift(item.msg.input.address);
                        message[Constant.TRANSACTIONMESSAGENAME.OUTPUTADDRESS].unshift(item.msg.output.address);
                        message[Constant.TRANSACTIONMESSAGENAME.INPUT].unshift(`${item.msg.input.coin.amount} ${item.msg.input.coin.denom.toUpperCase()}`);
                        message[Constant.TRANSACTIONMESSAGENAME.DEADLINE].unshift(item.msg.deadline);
                        message[Constant.TRANSACTIONMESSAGENAME.ISBUYORDER].unshift(item.msg.is_buy_order);
                        if(item.msg.output.coin.denom === Constant.Denom.IRISATTO){
                            amountObj =  Tools.formatAmountOfTxDetail(item.msg.output.coin);
                            message[Constant.TRANSACTIONMESSAGENAME.OUTPUT].unshift(amountObj);
                        }else {
                            message[Constant.TRANSACTIONMESSAGENAME.OUTPUT].unshift(`${item.msg.input.coin.amount} ${item.msg.input.coin.denom.toUpperCase()}`);
                        }
                    }
                }
            })
        }
        return message
    }
}
