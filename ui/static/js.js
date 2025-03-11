// Manage Servers

function onPageLoad() {
  const manageServerBtn = document.getElementById("serverListBtn");
  console.log(manageServerBtn);
  console.log("vagina");
}

function openServerDialog() {
  const dialog = document.getElementById("serverDialog");
  const dialogTmpl = document.getElementById("addServerDialogTemplate");
  //const dialogTmpl = document.getElementById("manageServerDialogTemplate");
  const dialogHtml = dialogTmpl.content.cloneNode(true);
  dialog.appendChild(dialogHtml);
  dialog.showModal();
}

function closeServerManagerModal() {
  const dialog = document.getElementById("serverDialog");
  dialog.close();
}

function openAddServerDialog() {
  const dialog = document.getElementById("serverDialog");
}

function goBackToServerManager() {
  const dialog = document.getElementById("serverDialog");
  const dialogTmpl = document.getElementById("manageServerDialogTemplate");
  dialog.replaceChildren(dialogTmpl.content.cloneNode(true));
}

function addServer() {
  let server = {};
  server.Id = -1;
  server.Name = document.getElementById("serverNameInput").value;
  server.Addr = document.getElementById("serverAddrInput").value;
  server.Port = document.getElementById("serverPortInput").value;
  server.Password = document.getElementById("serverPasswordInput").value;

  saveServer(server);
  console.log(loadServerList());
}

function loadServerList() {
  const serverListString = localStorage.getItem("serverlist");
  if (serverListString == null) {
    return [];
  } else {
    return JSON.parse(serverListString);
  }
}

function saveServer(server) {
  serverList = loadServerList();
  console.log(serverList.length);
  if (server.Id == -1) {
    server.Id = serverList.length;
    serverList.push(server);
  } else {
    serverList[server.Id].Name = server.Name;
    serverList[server.Id].Addr = server.Addr;
    serverList[server.Id].Port = server.Port;
    serverList[server.Id].Password = server.Password;
  }

  localStorage.setItem("serverlist", JSON.stringify(serverList));
}
