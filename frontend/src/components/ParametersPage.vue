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
    import Service from '../util/axios'
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
                    key:'test',
                    'Description/Usage' : '',
                    'Genesis Value' :'',
                    'Range' : '',
                    'Type' : '',
                    'Note' : ''
                  }]
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
                console.log(e)
              })

            },
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
