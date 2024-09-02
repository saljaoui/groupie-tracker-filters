let loadMore = document.querySelector('#load-more');
let loadLess = document.querySelector('#Load-less');
let btnfilter=document.querySelector('#filter-header');
let filter=document.querySelector('.filters-artist');

let clickee = 0;
btnfilter.addEventListener("click",()=>{
    if (clickee === 0) {
        filter.style.display='block'
        clickee = 1
    }
    else if (clickee === 1) {
        filter.style.display='none'
        clickee = 0
    }
  })

let currentItem = 24;

loadMore.onclick = () => {
    let boxes = [...document.querySelectorAll('.cards .card')];
    for (let i = currentItem; i >= currentItem - 24 && i < boxes.length ; i++){
        boxes[i].style.display = 'inline-block'
    }
    currentItem += 24;
    if (currentItem >= 48) {
        loadMore.style.display = 'none'
    }
    if (currentItem >= 28) {
        loadLess.style.display = 'flex'
    }

};

loadLess.onclick = () => {
    let boxes = [...document.querySelectorAll('.cards .card')];
    currentItem -= 24;
    for (let i = currentItem; i >= currentItem - 24 && i < boxes.length ; i++){
        boxes[i].style.display = 'none'
    }
    if (currentItem <= 48) {
        loadMore.style.display = 'flex'
    }
    console.log(currentItem)
    if (currentItem <= 24) {
        loadLess.style.display = 'none'
    }
};

document.getElementById('from-year').addEventListener('input', function() {
    document.getElementById('from-yearValue').textContent = this.value;
  });
  document.getElementById('to-year').addEventListener('input', function() {
    document.getElementById('to-yearValue').textContent = this.value;
  });

