let xhr = new XMLHttpRequest();
xhr.open('GET', '/api/get', true);
xhr.send();
xhr.onload = function () {
  if (xhr.readyState === 4 && xhr.status === 200) {
    let data = JSON.parse(xhr.responseText);
    renderTableWithFoldersAndDocs(data)
  }
}

function onFolderClick(objectId) {
  const xhr = new XMLHttpRequest();

  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
      let data = JSON.parse(xhr.responseText);
      renderTableWithFoldersAndDocs(data)
    }
  };
  xhr.open('GET', `/api/getId?folderId=${objectId}`, true);
  xhr.send();
}

// createHeadTableForFoldersAndDocs - создать шапку таблицы для файлов и документов.
function createHeadTableForFoldersAndDocs() {
  document.querySelector('tr').innerHTML = `
        <th>Тип</th>
        <th>Наименование</th>
        <th>Дата создания</th>
        <th>Дата изменения</th>
        <th>ФИО автора</th>
        <th>Подписано</th>
        <th>Тип подписи</th>
      `
}

// renderTableWithFoldersAndDocs - отобразить данные в таблице.
const renderTableWithFoldersAndDocs = data => {
  createHeadTableForFoldersAndDocs();
  const tableBody = document.querySelector('tbody');

  tableBody.innerHTML = '';

  if (data['Folders'] !== null) {
    data['Folders'].forEach(folder => {
      tableBody.innerHTML += `
      <tr class="folder" >
        <td onclick="onFolderClick(${folder['ObjectID']})"></td>
        <td>${folder['ObjectName']}</td>
        <td>${folder['Created']}</td>
        <td>${folder['Modified']}</td>
        <td>${folder['AuthorID']}</td>
        <td></td>
        <td></td>
      </tr>
    `
    })
  }

  if (data['Documents'] !== null) {
    data['Documents'].forEach(doc => {
      tableBody.innerHTML += `
      <tr class="file">
        <td onclick=onFileClick(${doc['ObjectID']})></td>
        <td>${doc['ObjectName']}</td>
        <td>${doc['Created']}</td>
        <td>${doc['Modified']}</td>
        <td>${doc['AuthorID']}</td>
        <td>${doc['Signed']}</td>
        <td>${doc['SignatureType']}</td>
      </tr>
    `
    })
  }
}

let idDocVer = 0;

const onFileClick = docId => {
  const xhr = new XMLHttpRequest();
  idDocVer = docId

  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
      console.log(xhr.responseText)
      let data = JSON.parse(xhr.responseText);
      renderTableWithDocVer(data)
    }
  };
  xhr.open('GET', `/api/getOnVer?docId=${docId}`, true);
  xhr.send();
};

const onOpenFileClick = (docId, verId) => {
  const xhr = new XMLHttpRequest();

  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
      console.log(xhr.responseType)
      var file = window.URL.createObjectURL(xhr.responseType);
      var a = document.querySelector("a");
      a.href = file;      window.open(file);
    }

      /*let data = JSON.parse(xhr.responseType);
      renderTableWithDocVer(data)*/
  };
  xhr.open('GET', `/api/getOnVer?docId=${docId}&verId=${verId}`, true);
  xhr.send();
};

// createHeadTableForFoldersAndDocs - создать шапку таблицы для документа с версией(ями).
function createHeadTableForDocWithVer() {
  document.querySelector('tr').innerHTML = `
        <th>Тип</th>
        <th>Наименование документа</th>
        <th>Версия файла</th>
        <th>Примечание</th>
        <th>Дата изменения</th>
        <th>Размер документа</th>
        <th>ФИО автора</th>
        <th>Наименование редактора</th>
        <th>Формат документа</th>
      `
}

// renderTableWithDocVer - Отобразить данные в таблице.
const renderTableWithDocVer = data => {
  createHeadTableForDocWithVer();
  const tableBody = document.querySelector('tbody');

  tableBody.innerHTML = '';

  data['DocumentWithVersions'].forEach(doc => {
    console.log(doc)

    tableBody.innerHTML += `
      <tr>
        <td onclick="onOpenFileClick(idDocVer, ${doc['VersionNumber']})"><a></a></td>
        <td>${doc['FileName']}</td>
        <td>${doc['VersionNumber']}</td>
        <td>${doc['Note']}</td>
        <td>${doc['ModifyDate']}</td>
        <td>${doc['Size']}</td>
        <td>${doc['AuthorDisplayFld']}</td>
        <td>${doc['EditorDisplayFld']}</td>
        <td>${doc['TypeVersionData']}</td>
      </tr>
                            `
  })
}