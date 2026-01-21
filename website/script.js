const ws = new WebSocket("ws://localhost:8765");
webID = null;

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  
  console.log("Received from server:", data);

  if (webID === null) {
    webID = data.array["webID"];
    console.log("Assigned webID:", webID);
  }
  
  document.getElementById("arrayDisplay").innerText = JSON.stringify(data.array);
  
};

function pushValue() {
  const XValue = document.getElementById("XValueInput").value;
  const YValue = document.getElementById("YValueInput").value;
  
  console.log("Pushing values:", XValue, YValue, webID);
  ws.send(JSON.stringify({
    action: "push",
    value: { XValue, YValue , webID }
  }));
}