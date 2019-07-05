let mouseenterEventFunc = () => {};
let mouseleaveEventFunc = () => {};

function bindHoverEvent (el, bind) {
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

function unbindHoverEvent (el) {
    if (el.removeEventListener) {
        el.removeEventListener('mouseenter', mouseenterEventFunc, false);
        el.removeEventListener('mouseleave', mouseleaveEventFunc, false);
    } else if (el.dettachEvent) {
        el.dettachEvent('onmouseenter', mouseenterEventFunc);
        el.dettachEvent('onmouseleave', mouseleaveEventFunc);
    }
}

function mouseleaveEvent (el) {
    return () => {
        let tooltipSpan = el.querySelector('.tooltip_span');
        tooltipSpan.style.opacity = '0';
        tooltipSpan.classList.remove('tooltip_span_word_warp');
        tooltipSpan.style.width = `auto`;
    }
}

function mouseenterEvent (el, containerDiv) {
    return () => {
        let containerRect = containerDiv.getBoundingClientRect();
        let elRect = el.getBoundingClientRect();

        let tooltipSpan = el.querySelector('.tooltip_span');
        let tooltipSpanI = tooltipSpan.querySelector('i');
        let tooltipSpanRect = tooltipSpan.getBoundingClientRect();

        let left = tooltipSpanRect.left - containerRect.left;
        let right = containerRect.right - tooltipSpanRect.right;

        let all = left + right;
        let elRectLeft = elRect.left - containerRect.left;
        let elRectRight = containerRect.right - elRect.right;
        tooltipSpan.style.opacity = '0';
        if (all < 0) {
            let maxWidth = Math.min(tooltipSpanRect.width, containerRect.width);
            tooltipSpan.classList.add('tooltip_span_word_warp');
            tooltipSpan.style.width = `${maxWidth}px`;
            let x = elRectLeft + elRect.width / 2;
            tooltipSpan.style.transform = `translateX(${-x}px)`;
            tooltipSpanI.style.left = `${x}px`;
        } else {
            if(right < 0) {
                let x = -tooltipSpanRect.width + elRectRight + elRect.width / 2;
                tooltipSpan.style.transform = `translateX(${x}px)`;
                tooltipSpanI.style.left = `${-x}px`;
            } else if (left < 0) {
                let x = elRectLeft + elRect.width / 2;
                tooltipSpan.style.transform = `translateX(${-x}px)`;
                tooltipSpanI.style.left = `${x}px`;
            }
        }
        tooltipSpan.style.opacity = '1';
    }
}

export default {
    name: "table-tooltip",
    hook: {
        inserted: function(el, bind) {
            if (bind.value.show) {
                bindHoverEvent(el, bind);
            }
        },
        unbind: function(el, bind) {
            if (bind.value.show) {
                unbindHoverEvent(el);
            }
        }
    }
};