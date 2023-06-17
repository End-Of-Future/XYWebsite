const url="http://localhost:2047/items";
function GetTitle(){
    var req=new XMLHttpRequest();
    req.open("POST",url,false);
    req.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    req.send("mode=title&code="+Math.random());
    document.getElementById("tit").innerText=req.responseText;
}
function GetTopItems(){
    var req=new XMLHttpRequest();
    req.open("POST",url,false);
    req.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    req.send("mode=t&code="+Math.random());
    const itArr = req.responseText.split(" ");
    var tmenu = document.getElementById("topm");
    for(it in itArr){
        tmenu.innerHTML+='<a>'+itArr[it]+'</a>';
    }
}
function GetLeftItems(path){
    var req=new XMLHttpRequest();
    req.open("POST",url,false);
    req.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    req.send("mode=l&path=" + path + "&code="+Math.random());
    const itArr = req.responseText.split(" ");
    var tmenu = document.getElementById("lmenu");
    for(it in itArr){
        tmenu.innerHTML+='<a>'+itArr[it]+'</a>';
    }
}