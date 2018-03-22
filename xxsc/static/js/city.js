var domain_url = 'http://' + document.domain + '/';

function getCity(e) {

    var pId = $(e).val();
    if (pId == 0)
        return;
    initDistrict();
    initCity();
    $.post(domain_url + "api/City/City", { proId: pId }, function (d) {
        var html =  "";
        $.each(d, function (i, val) {
            html += "<option value='" + val.Id + "'> " + val.Name + "</option>";
        });
        $("#CityId").append(html);
    }); 
}

function getDistrict(e){
    var cId = $(e).val();
    if (cId == 0)
        return;
    initDistrict();
    $.post(domain_url + "api/City/District", { cId: cId }, function (d) {
        var html = " ";
        $.each(d, function (i, val) {
            html += "<option value='" + val.Id + "'> " + val.Name + "</option>";
        });
        $("#DistrictId").append(html);
    });
}



function initCity() {
    $("#CityId").children().remove();
    var html = " <option value='0'>地市</option>";
    $("#CityId").html(html);
}

function initDistrict() {
    $("#DistrictId").children().remove();
    var html = " <option value='0'>县区</option>";
    $("#DistrictId").html(html);
}


