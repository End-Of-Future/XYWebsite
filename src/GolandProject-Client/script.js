const url="http://localhost:2047/items";
function GetTitle(){
    let req=new XMLHttpRequest();
    req.open("POST",url,false);
    req.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    req.send("mode=title&code="+Math.random());
    document.getElementById("tit").innerText=req.responseText;
}
function GetTopItems(){
    let req=new XMLHttpRequest();
    req.open("POST",url,false);
    req.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    req.send("mode=t&code="+Math.random());
    const itArr = req.responseText.split(" ");
    let tmenu = document.getElementById("topm");
    for(let i=0;i<itArr.length;++i){
        let it=itArr[i]
        const ik = it.split("=");
        tmenu.innerHTML+='<a href=\"'+ik[1]+ '\">'+ik[0]+'</a>';
    }
}
function GetLeftItems(path){
    let req=new XMLHttpRequest();
    req.open("POST",url,false);
    req.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    req.send("mode=l&path=" + path + "&code="+Math.random());
    const itArr = req.responseText.split(" ");
    let tmenu = document.getElementById("lmenu");
    for(let i=0;i<itArr.length;++i){
        let it=itArr[i]
        const ik = it.split("=");
        tmenu.innerHTML+='<a href=\"'+ik[1]+ '\">'+ik[0]+'</a>';
    }
}
function GetContent(path){
    let req=new XMLHttpRequest();
    req.open("POST",url,false);
    req.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    req.send("mode=c&path=" + path + "&code="+Math.random());
    document.getElementById("main").innerText=req.responseText;
}