// Поисковик.
window.onload = () => {
  document.querySelector('#search-text').oninput = function () {
    let value = this.value.trim();
    let list = document.querySelectorAll('.choice');

    if (value !== '') {
      list.forEach(elem => {
        if (elem.innerHTML.search(value) === -1) {
          elem.classList.add('hide');
        } else {
          elem.classList.remove('hide');
        }
      });
    } else {
      list.forEach(elem => {
        elem.classList.remove('hide');
      })
    }
  }
}