import AmountFormat from "./AmountFormat";

export default {
	install(Vue) {
        Vue.filter(AmountFormat.name, AmountFormat.filter);
	}
}
