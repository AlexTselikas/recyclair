
var mapOptions = {
  center: [35.3180305,25.1018764],
  zoom:15
}
var map = new L.map('map',mapOptions);
var layer = new L.TileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',{attribution:'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors',maxZoom:20});
map.addLayer(layer);

function onMapClick(e) {
  popup
      .setLatLng(e.latlng)
      .setContent("You clicked the map at <br/>" + e.latlng.toString())
      .openOn(map);
}

map.on('click', onMapClick);

var garbageBinInfo = {
    title: "Garbage Bin",
    icon:L.icon({iconUrl:'http://ermescloud.net/garbage-bin.png',iconSize:[35,40]})
}
var recyclingBinInfo = {
    title: "Recycling Bin",
    icon:L.icon({iconUrl:'http://ermescloud.net/recycling-bin.png',iconSize:[35,40]})
}
var bothInfo = {
    title: "Recycling and Garbage Bin",
    icon:L.icon({iconUrl:'http://ermescloud.net/both.png',iconSize:[35,40]})
}
var marker = L.marker([35.316201,25.101393],bothInfo);
marker.bindPopup("Recycling and Garbage Bin");
marker.addTo(map);
var marker = L.marker([35.317013,25.101373],garbageBinInfo);
marker.bindPopup("Garbage Bin");
marker.addTo(map);
var marker = L.marker([35.318307 ,25.101186],garbageBinInfo);
marker.bindPopup("Garbage Bin");
marker.addTo(map);
var marker = L.marker([35.320021 ,25.102349],garbageBinInfo);
marker.bindPopup("Garbage Bin");
marker.addTo(map);
var marker = L.marker([35.319647 ,25.101875],bothInfo);
marker.bindPopup("Recycling and Garbage Bin");
marker.addTo(map);
