function getimage(){
    beatids = Array.from(document.querySelectorAll('[id^="beat"]'));
    beatids.sort();
    console.log(beatids)
    values = [];
    beatids.forEach(b => values.push(document.querySelector(b.id).value));
    console.log(values);
    
}