import TableHeadFixed from './TableHeadFixed';

export default {
	install(Vue) {
		Vue.directive(TableHeadFixed.name, TableHeadFixed.hook)
	}
}