function searchArea() {
    let zip_cd = document.getElementById('zip_cd').value;
    if (zip_cd === "") {
        alert("郵便番号を入力してください。");
        return;
    }
    var req = new XMLHttpRequest();
    req.onreadystatechange = function() {
        let result = document.getElementById('area');
        if (req.readyState == 4) { // 通信の完了時
          if (req.status == 200) { // 通信の成功時
            //result.innerHTML = req.responseText;
            var html = "";
            let areas = JSON.parse(req.responseText);
            for (let i = 0; i < areas.length; i++) {
                let area = areas[i];
                html += "<p>"
                html += area.pref_name + area.city_name + area.town_name;
                html += "</p>"

                let elem = document.getElementById('notify_url');
                elem.href = elem.href + area.id;
            }
            result.innerHTML = html;

          }
        }else{
          result.innerHTML = "通信中...";
        }
    }
    req.open('GET', '/gc_alert/search_area?zip_cd=' + zip_cd, true);
    req.send(null);
}

function checkLineNotifyUrl() {
    let url = document.getElementById('notify_url').href;
    if (url.endsWith('=')) {
        alert('地域検索を行ってください。');
        return false;
    }
    return true;
}