<template>
    <div>
        <m-table :data="items" :columns="fields">
            <template slot-scope="{ row }" slot="moniker">
                <div class="moniker_conent">
                    <div class="proposer_img_content">
                        <img :style="{visibility:row.flProposer ? 'visible' : 'hidden'}" src="../../assets/proposer_img.png"/>
                    </div>
                    <span class="skip_route">
                        <router-link :to="addressRoute(row.moniker)">{{row.moniker? row.moniker :''}}</router-link>
                    </span>
                </div>
            </template>
            <template slot-scope="{ row }" slot="OperatorAddress">
                <div class="common_hover_address_parent skip_route">
                    <router-link :to="addressRoute(row.OperatorAddress)" class="link_style common_font_style">{{formatAddress(row.OperatorAddress)}}
                    </router-link>
                </div>
            </template>
        </m-table>
    </div>
</template>

<script>
    import Tools from "../../util/Tools"
	export default {
		name: "MBlocKInformationTable",
        data() {
			return {
				fields:[
                    {
	                    title: '#',
	                    key: 'index',
                    },
					{
						title:'Moniker',
						slot:'moniker',
					},
					{
						title:'Operator',
						slot:'OperatorAddress',
						tooltip: true
					},
					{
						title:'Consensus Address',
						key:'Consensus'
					},
					{
						title:'Proposer Priority',
						key:'ProposerPriority'
					},
					{
						title:'Voting Power',
						key:'VotingPower'
					}
                ]
            }
        },
		props:{
	        items:{
		        type: Array,
		        default: []
            }
        },
		watch:{
			items(items){
				if(items){
					this.items.map((item,i)=>{
						return item.index = i + 1
					})
				}
				this.items = this.items;
			}
		},
        methods:{
	        formatAddress (address) {
		        if (address) {
			        return Tools.formatValidatorAddress(address)
		        }
	        },
        }

	}
</script>

<style lang="scss">
    .block_validator_set_content{
        .m-table-header{
            table{
                thead{
                    tr{
                        th:nth-of-type(2){
                            padding-left: 0.27rem;
                        }
                    }
                }
            }
        }
    }
</style>
