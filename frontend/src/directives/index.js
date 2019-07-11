import TableHeadFixed from './TableHeadFixed';
import TableTooltip from './TableTooltip';

export default {
	install(Vue) {
        Vue.directive(TableHeadFixed.name, TableHeadFixed.hook);
        Vue.directive(TableTooltip.name, TableTooltip.hook);
	}
}
