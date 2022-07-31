function getImage(){
    query = getImageQuery();
    fetch(query).then((resp )=> resp.json()).then((respJSON) => console.log(respJSON));
}

function getImageQuery() {
    beatids = Array.from(document.querySelectorAll('[id^="beat"]'));
    beatids.sort();
    values = [];
    beatids.forEach(b => values.push(document.getElementById(b.id).value));
    // remove empties
    values.filter(v => v.length > 0)
    // build the query string
    base = "http://localhost:8050/GetMeasureImage?";
    params = [];
    values.forEach((v, i) => params.push(`b${i}=${v}`));
    qparams = params.join('&');
    query = base + qparams;
    return query
}