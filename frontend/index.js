const tvshowSaveForm = document.getElementById("tvshowSaveForm");
const tvshowFindForm = document.getElementById("tvshowFindForm");
const tvshowDataField = document.getElementById("tvshowDataField");

function tvshowSaveHandler(event) {
    event.preventDefault();

    const tvshowName = document.getElementById('tvshowName1').value;
    console.log(tvshowName)

    fetch('/api/save', {
        method: 'POST',
        body: JSON.stringify({
            'name': tvshowName,
        }),
        headers: {
            'Content-Type': 'application/json',
        },
    }).then(response => {
        response.json().then((res) => {
            const tvshowActionField = document.getElementById("save-action");
            const tvshowStatusField = document.getElementById("save-status");
            console.log(res);
            tvshowActionField.innerHTML = res.data;
            tvshowStatusField.innerHTML = res.status;
        })
    }).catch(err => console.log("Error when /save", err));
}

function tvshowFindHandler(event) {
    event.preventDefault();

    const tvshowName = document.getElementById('tvshowName2').value;

    fetch('/api/find/' + tvshowName).then(response => {
        response.json().then((res) => {
            const tvshowIdField = document.getElementById("find-id");
            const tvshowNameField = document.getElementById("find-name");
            console.log(res);
            tvshowIdField.innerHTML = res._id;
            tvshowNameField.innerHTML = res.name;
        })
    }).catch(err => console.log("Error when /find", err));
}

tvshowSaveForm.addEventListener('submit', tvshowSaveHandler);
tvshowFindForm.addEventListener('submit', tvshowFindHandler);