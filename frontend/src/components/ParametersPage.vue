<template>
  <div class="parameters_page_container">
    <div class="parameters_page_title_container">
      <div class="parameters_title_content">
        <span class="parameters_title">Parameters</span>
      </div>
    </div>
    <div class="parameters_list_container">
      <div class="parameter_list_content">
        <h5 class="parameter_list_title_content">All Governance Params In IRISnet</h5>
        <div class="parameters_list_table_wrap">
          <div class="parameters_list_table_content">
            <spin-component :showLoading="showLoading"></spin-component>
            <v-parameters :items="parametersList" :showNoData="showNoData"></v-parameters>
            <div v-show="showNoData" class="no_data_show">
              No Data
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
    import VParameters from "./table/ParametersTable";
    import SpinComponent from './commonComponents/SpinComponent';
    import Service from '../util/axios';
    import Tools from "../util/Tools"
    export default {
        name: "Parameters",
        components: {SpinComponent, VParameters},
        data() {
            return {
              parametersList:[],
              proposalsListPageWrap: '',
              showLoading:false,
              showNoData: false,
            }
        },
        mounted() {
          this.getParametersList()
        },
        methods:{
            getParametersList(){
              this.showLoading = true;
              let parametersListUrl = `/api/gov/params`;
              Service.http(parametersListUrl).then((res) => {
                this.showLoading = false;

                if(res){
                  this.parametersList = res.map( item => {
                    this.handleParameterItem(item);
                    return {
                      key: `${item.module}/${item.key}`,
                      'Description/Usage' : item.description,
                      'Genesis Value' : item.value,
                      'Range' : item.range,
                      'Type' : item.type,
                      'Note' : item.note
                    }
                  })
                }else {
                  this.parametersList = [{
                    key:'',
                    'Description/Usage' : '',
                    'Genesis Value' :'',
                    'Range' : '',
                    'Type' : '',
                    'Note' : ''
                  }];
                  this.showNoData = true;
                }
              }).catch(e => {
                this.showLoading = false;

                this.parametersList = [{
                  key:'',
                  'Description/Usage' : '',
                  'Genesis Value' :'',
                  'Range' : '',
                  'Type' : '',
                  'Note' : ''
                }];
                this.showNoData = true;
                console.error(e)
              })

            },
            handleParameterItem(parameterItem){
              if(parameterItem.key === "max_validators"){
              }else if(parameterItem.key === "unbonding_time"){
                parameterItem.value = this.formatUnbondingTime(parameterItem.value);
                parameterItem.range.minimum.data = this.formatUnbondingTime(parameterItem.range.minimum.data)
                parameterItem.range.maximum.data = this.formatUnbondingTime(parameterItem.range.maximum.data)
              }else if(parameterItem.key === "inflation"){
                parameterItem.value = `${Number(parameterItem.value)} %`;
                parameterItem.range.minimum.data = `${Number(parameterItem.range.minimum.data)} %`;
                parameterItem.range.maximum.data = `${Number(parameterItem.range.maximum.data)} %`
              }else if(parameterItem.key === "base_proposer_reward"){
                parameterItem.value = `${Number(parameterItem.value)} %`;
                parameterItem.range.minimum.data = `${Number(parameterItem.range.minimum.data)} %`;
                parameterItem.range.maximum.data = `${Number(parameterItem.range.maximum.data)} %`
              }else if(parameterItem.key === "bonus_proposer_reward"){
                parameterItem.value = `${Number(parameterItem.value)} %`;
                parameterItem.range.minimum.data = `${Number(parameterItem.range.minimum.data)} %`;
                parameterItem.range.maximum.data = `${Number(parameterItem.range.maximum.data)} %`
              }else if(parameterItem.key === "community_tax"){
                parameterItem.value = `${Number(parameterItem.value)} %`;
                parameterItem.range.minimum.data = `${Number(parameterItem.range.minimum.data)} %`;
                parameterItem.range.maximum.data = `${Number(parameterItem.range.maximum.data)} %`
              }else if(parameterItem.key === "gas_price_threshold"){
                parameterItem.value = `${Tools.formatNumber(parameterItem.value)} IRIS`;
                parameterItem.range.maximum.data = `${Tools.formatNumber(parameterItem.range.maximum.data)} IRIS`
              }else if(parameterItem.key === "tx_size"){
                parameterItem.value = `${parameterItem.value} Byte`;
                parameterItem.range.minimum.data = `${parameterItem.range.minimum.data} Byte`;
                parameterItem.range.maximum.data = `${parameterItem.range.maximum.data} Byte`
              }
            },
            formatUnbondingTime(time) {
                let nsToMSRatio = 1000000, dToHRatio = 24, HToMRatio = 60;
                let dateTime = Tools.formatDuring(Number(time) / nsToMSRatio), d, h, m;
                if (dateTime.days > 1) {
                    d = `${Math.floor(dateTime.days)}d`
                } else if (dateTime.days === 1) {
                    d = `${Math.floor(dateTime.days)}d`
                } else {
                    d = ''
                }
                if (dateTime.hours > 1 && dateTime.hours < dToHRatio) {
                    h = `${Math.floor(dateTime.hours)}h`
                } else if (dateTime.hours === 1) {
                    h = `${Math.floor(dateTime.hours)}h`
                } else {
                    h = ''
                }
                if (dateTime.minutes > 1 && dateTime.minutes < HToMRatio) {
                    m = `${Math.floor(dateTime.minutes)}m`
                } else if (dateTime.minutes === 1) {
                    m = `${Math.floor(dateTime.minutes)}m`
                } else {
                    m = ''
                }
                return `${d ? d :''}${h ? h :''}${m ? m :''}`
            }
        }
    }
</script>

<style scoped lang="scss">
@import "../style/mixin.scss";
.parameters_page_container{
  .parameters_page_title_container{
    background: rgba(239, 239, 241, 1);
    border-bottom: 0.01rem solid rgba(215, 217, 224, 1);
    .parameters_title_content{
      max-width: 12.8rem;
      margin: 0 auto;
      .parameters_title{
        height: 0.6rem;
        display: flex;
        align-items: center;
        font-size: 0.22rem;
        color: rgba(34, 37, 42, 1);
        line-height: 0.26rem;
        padding-left: 0.2rem;
      }
    }
  }
  .parameters_list_container{
    width: 100%;
    margin-bottom: 0.2rem;
    .parameter_list_content{
      max-width: 12.8rem;
      width: 100%;
      margin: 0 auto;
      .parameter_list_title_content{
        font-size: 0.18rem;
        height: 0.7rem;
        display: flex;
        align-items: center;
        padding-left: 0.2rem;
        line-height: 0.25rem;
      }
      .parameters_list_table_wrap{
        overflow-x: auto;
        .parameters_list_table_content{
          min-width: 11.5rem;
        }
      }
    }
  }
}
.no_data_show {
  @include flex;
  justify-content: center;
  border-top: 0.01rem solid #eee;
  border-bottom: 0.01rem solid #eee;
  font-size: 0.14rem;
  height: 3rem;
  align-items: center;
}
</style>
