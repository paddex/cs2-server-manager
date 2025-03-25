// Manage Servers

function onPageLoad() {
  updateMainServerList();
}

async function updateMainServerList() {
  const servers = loadServerList();
  const response = await fetch("/serverinfo", {
    method: "POST",
    body: JSON.stringify(servers),
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (response.status != 200) {
    console.log("Something went wrong!");
    return;
  }
  const list = await response.json();

  const table = document
    .querySelector("#mainTableTemplate")
    .content.cloneNode(true);
  let tbody = table.querySelector("tbody");
  for (l of list) {
    const row = document
      .querySelector("#mainServerRowInfoTemplate")
      .content.cloneNode(true);
    let tds = row.querySelectorAll("td");
    tds[0].textContent = l.Name;
    tds[1].textContent = l.Addr;
    tds[2].textContent = l.Port;
    tds[3].textContent = l.Map;
    tds[4].textContent = l.Players;

    tbody.appendChild(row);
  }
  const tablediv = document.querySelector("#mainTableDiv");
  const spinner = tablediv.querySelector("svg");
  if (spinner != undefined) {
    tablediv.removeChild(spinner);
  }
  const oldTable = tablediv.querySelector("#mainTable");
  if (oldTable != undefined) {
    tablediv.removeChild(oldTable);
  }
  tablediv.appendChild(table);
}

function openServerDialog() {
  const dialog = document.getElementById("serverDialog");
  loadServerDialogServers();
  dialog.showModal();
}

function closeServerManagerModal() {
  const dialog = document.getElementById("serverDialog");
  dialog.close();
}

function openAddServerDialog(server = null) {
  const dialog = document.getElementById("serverDialog");
  const dialogContent = document
    .getElementById("addServerDialogTemplate")
    .content.cloneNode(true);
  if (server != null) {
    dialogContent.querySelector("#idInput").value = server.Id;
    dialogContent.querySelector("#serverNameInput").value = server.Name;
    dialogContent.querySelector("#serverAddrInput").value = server.Addr;
    dialogContent.querySelector("#serverPortInput").value = server.Port;
    dialogContent.querySelector("#serverPasswordInput").value = server.Password;
  }
  dialog.replaceChildren(dialogContent);
}

function goBackToServerManager() {
  loadServerDialogServers();
}

function addServer() {
  const id = document.getElementById("idInput").value;
  let server = {};
  server.Name = document.getElementById("serverNameInput").value;
  server.Addr = document.getElementById("serverAddrInput").value;
  server.Port = parseInt(document.getElementById("serverPortInput").value, 10);
  server.Password = document.getElementById("serverPasswordInput").value;

  saveServer(server, id);
  goBackToServerManager();
}

function loadServerList() {
  const serverListString = localStorage.getItem("serverlist");
  if (serverListString == null) {
    return [];
  } else {
    return JSON.parse(serverListString);
  }
}

function saveServerList(list) {
  localStorage.setItem("serverlist", JSON.stringify(list));
}

function saveServer(server, i = -1) {
  serverList = loadServerList();
  if (i == -1) {
    serverList.push(server);
  } else {
    serverList[i].Name = server.Name;
    serverList[i].Addr = server.Addr;
    serverList[i].Port = server.Port;
    serverList[i].Password = server.Password;
  }

  saveServerList(serverList);
}

function loadServerDialogServers() {
  serverList = loadServerList();
  const dialogTmpl = document.getElementById("manageServerDialogTemplate");
  const dialogContent = dialogTmpl.content.cloneNode(true);
  const rowTmpl = document.getElementById("serverRowTemplate");

  for (const [i, server] of serverList.entries()) {
    const row = rowTmpl.content.cloneNode(true);
    tds = row.querySelectorAll("td");
    tds[0].textContent = server.Name;
    tds[1].textContent = server.Addr;
    tds[2].textContent = server.Port;
    tds[3].textContent = server.Password;

    tr = row.querySelector("tr");
    tr.dataset.id = i;
    tr.dataset.name = server.Name;
    tr.dataset.addr = server.Addr;
    tr.dataset.port = server.Port;
    tr.dataset.password = server.Password;

    dialogContent.querySelector("tbody").appendChild(row);
  }

  const dialog = document.getElementById("serverDialog");
  dialog.replaceChildren(dialogContent);
}

function setServerManagerRowActive(tr) {
  const id = tr.dataset.id;
  let currentId = -1;
  const activeRow = document.querySelector(".serverManagerActiveRow");
  if (activeRow != null) {
    activeRow.classList.remove("serverManagerActiveRow");
    currentId = activeRow.dataset.id;
  }
  if (id != currentId) {
    const allRows = document.querySelectorAll(".serverRow");
    allRows[id].classList.add("serverManagerActiveRow");
  }
}

function deleteServerFromServerList() {
  const i = document.querySelector(".serverManagerActiveRow").dataset.id;
  serverList = loadServerList();
  serverList.splice(i, 1);
  saveServerList(serverList);
  loadServerDialogServers();
}

function editServerDialog() {
  const row = document.querySelector(".serverManagerActiveRow");
  let server = {};
  server.Id = row.dataset.id;
  server.Name = row.dataset.name;
  server.Addr = row.dataset.addr;
  server.Port = row.dataset.port;
  server.Password = row.dataset.password;

  openAddServerDialog(server);
}
