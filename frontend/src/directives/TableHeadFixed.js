import Tools from '../util/Tools';
let resizeFn, resizeTime;

const changeScroll = (e) => {
  let element = document.querySelector('.m-table-header');
  if (element) {
    if (e.target.scrollLeft > 0) {
      element.style.left = -e.target.scrollLeft + 'px';
    } else {
      element.style.left = 'auto';
    }
  }
}

const windowRisize = (node) => {
  let element = document.querySelector('.m-table-body');
  if (element && node) {
    if (node.offsetWidth === element.offsetWidth) {
      document.querySelector('.m-table-header').style.left = 'auto';
    }
  }
}

const bundEvents = (el, vNode) => {
  let node = el && el.parentNode && el.parentNode.parentNode;
  if (!node) { return };
  resizeFn = () => {
    clearTimeout(resizeTime);
    resizeTime = setTimeout(() => {
      windowRisize(node);
      classFunc(el, vNode);
    }, 500);
  };
  if (node.addEventListener) {
    node.addEventListener('scroll', changeScroll, false);
    window.addEventListener('resize', resizeFn, false);
  } else if (node.attachEvent) {
    node.attachEvent('onscroll', changeScroll);
    window.attachEvent('onresize', resizeFn);
  }
}

const unbundEvents = (el) => {
  let node = el && el.parentNode && el.parentNode.parentNode;
  if (!node) { return };
  if (node.removeEventListener) {
    node.removeEventListener('scroll', changeScroll, false);
    window.removeEventListener('resize', resizeFn, false);
  } else if (node.dettachEvent) {
    node.dettachEvent('onscroll', changeScroll);
    window.dettachEvent('onresize', resizeFn);
  }
}

const classFunc = (el, vNode) => {
  if (Tools.currentDeviceIsPersonComputer() && !vNode.componentInstance.$store.state.isMobile) {
    el.classList.add('override_mtable');
  } else {
    el.classList.remove('override_mtable');
  }
}

export default {
  name: 'table-head-fixed',
  hook: {
    inserted: function(el, bind, vNode) {
      classFunc(el, vNode);
      bundEvents(el, vNode);
    },
    unbind: function(el) {
      unbundEvents(el);
    }
  }
}