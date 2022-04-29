/**
 * @license
 * Copyright 2019 Google LLC. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
function initMap(): void {
  const directionsService = new google.maps.DirectionsService();
  const directionsRenderer = new google.maps.DirectionsRenderer();
  const map = new google.maps.Map(
    document.getElementById("map") as HTMLElement,
    {
      zoom: 7,
      center: { lat: 41.85, lng: -87.65 },
    }
  );

  directionsRenderer.setMap(map);
  (document.getElementById("search") as HTMLElement).addEventListener("click", function ()
      {
          let waypts = [{location:(document.getElementById("origem") as HTMLInputElement).value,stopover: false}];
          let carona = {
              "Origem": (document.getElementById("origem") as HTMLInputElement).value,
              "Destino": (document.getElementById("destino") as HTMLInputElement).value
          };
          fetch("http://localhost:8080/search", {
              method: "POST",
              body: JSON.stringify(carona)
          }).then((response) => {
              response.text().then(async function (carona) {
                  let result = JSON.parse(carona);
                  directionsService
                      .route({
                          origin: {
                              query: result.Origem,
                          },
                          destination: {
                              query: result.Destino,
                          },
                          travelMode: google.maps.TravelMode.DRIVING,
                          waypoints: waypts,
                          optimizeWaypoints: true,
                      })
                      .then((response) => {
                          directionsRenderer.setDirections(response);
                          const route = response.routes[0];
                          const summaryPanel = document.getElementById(
                              "Sumary"
                          ) as HTMLElement;

                          summaryPanel.innerHTML = "";

                          // For each route, display summary information.
                          for (let i = 0; i < route.legs.length; i++) {
                              const routeSegment = i + 1;

                              summaryPanel.innerHTML +=
                                  "<b>Route Segment: " + routeSegment + "</b><br>";
                              summaryPanel.innerHTML += route.legs[i].start_address + " to ";
                              summaryPanel.innerHTML += route.legs[i].end_address + "<br>";
                              summaryPanel.innerHTML += route.legs[i].distance!.text + "<br><br>";
                          }
                      })
                      .catch((e) => window.alert("Directions request failed due to " + status));
              });
          }).catch((error) => {
              console.log(error)
          });
      }

  );

}

declare global {
  interface Window {
    initMap: () => void;
  }
}
window.initMap = initMap;
export {};
