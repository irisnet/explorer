let mouseenterEventFunc = () => { };
let mouseleaveEventFunc = () => { };

function bindHoverEvent(el, bind) {
    let containerDiv = bind.value.container.parentNode.parentNode;
    mouseenterEventFunc = mouseenterEvent(el, containerDiv);
    mouseleaveEventFunc = mouseleaveEvent(el, containerDiv);
    if (el.addEventListener) {
        el.addEventListener('mouseenter', mouseenterEventFunc, false);
        el.addEventListener('mouseleave', mouseleaveEventFunc, false);
    } else if (el.attachEvent) {
        el.attachEvent('onmouseenter', mouseenterEventFunc);
        el.attachEvent('onmouseleave', mouseleaveEventFunc);
    }
}

function unbindHoverEvent(el) {
    if (el.removeEventListener) {
        el.removeEventListener('mouseenter', mouseenterEventFunc, false);
        el.removeEventListener('mouseleave', mouseleaveEventFunc, false);
    } else if (el.dettachEvent) {
        el.dettachEvent('onmouseenter', mouseenterEventFunc);
        el.dettachEvent('onmouseleave', mouseleaveEventFunc);
    }
}

function mouseleaveEvent(el) {
    return () => {
        let tooltipSpan = el.querySelector('.tooltip_span');
        if (!tooltipSpan) { return }
        let tooltipSpanI = tooltipSpan.querySelector('i');
        tooltipSpan.classList.remove('tooltip_span_word_warp');
        tooltipSpan.removeAttribute("style");
        tooltipSpanI.removeAttribute("style");
    }
}

function mouseenterEvent(el, containerDiv) {
    return () => {
        let containerRect = containerDiv.getBoundingClientRect();
        let elRect = el.getBoundingClientRect();

        let tooltipSpan = el.querySelector('.tooltip_span');
        if (!tooltipSpan) { return }
        let tooltipSpanI = tooltipSpan.querySelector('i');
        let tooltipSpanRect = tooltipSpan.getBoundingClientRect();

        tooltipSpan.style.opacity = '0';
        tooltipSpan.style.left = `${elRect.left - (tooltipSpanRect.width - elRect.width) / 2}px`;

        tooltipSpanRect = tooltipSpan.getBoundingClientRect();

        let left = tooltipSpanRect.left - containerRect.left;
        let right = containerRect.right - tooltipSpanRect.right;
        let all = left + right;

        if (all < 0) {
            let maxWidth = Math.min(tooltipSpanRect.width, containerRect.width);
            tooltipSpan.classList.add('tooltip_span_word_warp');
            tooltipSpan.style.width = `${maxWidth}px`;

            tooltipSpanRect = tooltipSpan.getBoundingClientRect();
            tooltipSpan.style.top = `${elRect.top - tooltipSpanRect.height - 4}px`;
            left = tooltipSpanRect.left - containerRect.left;
            right = containerRect.right - tooltipSpanRect.right;
        } else {
            tooltipSpanRect = tooltipSpan.getBoundingClientRect();
            tooltipSpan.style.top = `${elRect.top - tooltipSpanRect.height - 4}px`;

            left = tooltipSpanRect.left - containerRect.left;
            right = containerRect.right - tooltipSpanRect.right;
        }
        if (right < 0) {
            tooltipSpan.style.transform = `translateX(${right}px)`;
        } else if (left < 0) {
            tooltipSpan.style.transform = `translateX(${-left}px)`;
        }
        tooltipSpanRect = tooltipSpan.getBoundingClientRect();
        let iLeft = elRect.left - tooltipSpanRect.left + elRect.width / 2;
        iLeft = Math.min(iLeft, tooltipSpanRect.width - 8);
        tooltipSpanI.style.transform = `translateX(${iLeft}px)`;
        tooltipSpan.style.opacity = '1';
    }
}

export default {
    name: "table-tooltip",
    hook: {
        inserted: function (el, bind) {
            if (bind.value.show) {
                bindHoverEvent(el, bind);
            }
        },
        unbind: function (el, bind) {
            if (bind.value.show) {
                unbindHoverEvent(el);
            }
        }
    }
};
