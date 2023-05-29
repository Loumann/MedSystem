const errorElement = document.querySelector("#container-notify");

const initAutoSearchUsers = () => {
	const searchElement = document.querySelector("#searchInput");
	if (!searchElement) {
		return;
	}

	searchElement.addEventListener('input', loadUsers);
};

class Fetch {
	baseURL = "";

	constructor(url) {
		this.baseURL = url;
	}

	signIn = (username, password, callback) => {
		this.POST("/sign-in", {
			username, password,
		}, callback)
	};

	createUser = (name, surname, patronymic, snils, callback) => {
		this.POST("/create-user", {
			name, surname, patronymic, snils,
		}, callback)
	};

	deleteUser = (id, callback) => {
		this.DELETE(`/delete-user/${id}`,  callback)
	};

	appAnalisys = (id, callback) =>
	{
		this.POST(`/analysis/${id}`, callback)
	}

	searchUsers = (search, callback) => {
		this.GET(`/users/${search}`, callback);
	};

	getAnalysis = (user, callback) => {
		this.POST(`/analysis/${user}`, callback);
	};

	GET = (uri, callback) => {
		const xhr = new XMLHttpRequest();
		xhr.open("GET", uri);
		xhr.onload = () => {

			if (xhr.status === 403 ) {
				alert("403 статус");
				return;
			}
			if (xhr.response) {
				callback(JSON.parse(xhr.response));
				return;
			}
			callback();
		};
		xhr.send();
	};

	POST = (uri, params, callback) => {
		const xhr = new XMLHttpRequest();
		xhr.open("POST", uri);
		xhr.onload = () => {
			if (xhr.status === 500 ) {
				showError("Неверный пароль\логин")
				return;
			}
			if (xhr.response) {
				errorElement.textContent = '';
				callback && callback(JSON.parse(xhr.response));
				return;
			}
			callback && callback();
		};
		xhr.send(JSON.stringify(params));
	};


	DELETE = (uri, callback) => {
		const xhr = new XMLHttpRequest();
		xhr.open("DELETE", uri);
		xhr.onload = () => {
			if (xhr.status !== 200) {
				// TODO ай-яй-яй
			}
			callback();

		};
		xhr.send();
	};
}
const fetch = new Fetch("http://localhost:8080");

function onSignIn() {
	const usernameElement = document.querySelector("#usernameInput");
	const passwordElement = document.querySelector("#passwordInput");

	const username = usernameElement.value.trim();
	const password = passwordElement.value.trim();


	if (username === '' || password === '' )
	{
		showError("Необходимо заполнить все поля")
		return;
	}

	fetch.signIn(username, password, () => {
		window.location.href = "/index";
		loadUsers();
	});
}

function loadUsers() {
	const searchElement = document.querySelector("#searchInput");
	if (!searchElement) {
		return;
	}

	const search = searchElement.value;
	fetch.searchUsers(search, (users) => {
		clearUsersList();
		fillUsersList(users);
	});
}

function onChangeUser() {
	const selectUserElement = document.querySelector("#userSelect");
	if (!selectUserElement) {
		return;
	}

	if (!selectUserElement.value) {
		clearAnalysesList();
		return;
	}

	fetch.getAnalysis(selectUserElement.value, (analyses) => {
		clearAnalysesList();
		fillAnalysesList(analyses);
	});
}

function onDeleteUser()
{
	const selectUserElement = document.querySelector("#userSelect")
	fetch.deleteUser(selectUserElement.value, () => {showSuccess("Пользователь удален", loadUsers())})
}

function appAnalisys()
{
	const selectUserElement = document.querySelector("#userSelect")
	fetch.appAnalisys(selectUserElement.value)
}

const clearAnalysesList = () => {
	const analiseBlock = document.querySelector("#AnaliseBlock");
	if (!analiseBlock) {
		return;
	}
	analiseBlock.innerHTML = "";
}

const fillAnalysesList = (analyses) => {
	const analiseBlock = document.querySelector("#AnaliseBlock");

	for (const analise of analyses) {
		const box = getAnaliseBox(analise);
		analiseBlock.append(box);
	}
};

const clearUsersList = () => {
	const selectUserElement = document.querySelector("#userSelect");
	selectUserElement.innerHTML = "";
};

const fillUsersList = (users) => {
	const selectUserElement = document.querySelector("#userSelect");

	for (const user of users) {
		const option =
			new Option(`${user.first_name} ${user.last_name} ${user.patronymic} ${user.snils}`, user.id);
		selectUserElement.options.add(option);
		selectUserElement.append(option);
	}

	selectUserElement.selectedIndex = 0;
	onChangeUser();
	selectUserElement.onchange = onChangeUser;
};

const getAnaliseBox = (analise) => {
	const box = document.createElement("div");
	box.className = "analise-box";

	box.append(getAnaliseRow("Дата и время сдачи", analise.date),
		getAnaliseRow("Эритроциты(Bld)", analise.bld),
		getAnaliseRow("Уробилиноген(UBG)", analise.ubg, (v) => v >= 0 && v <= 6 || v === undefined),
		getAnaliseRow("Билирубин(Bil)", analise.bil, (v) => v >= 1 && v <= 3 || v === undefined),
		getAnaliseRow("Белок(Pro)", analise.pro, (v) => v >= 0 && v <= 0.3 || v === undefined),
		getAnaliseRow("Нитриты(Nit)", analise.nit, (v) => v === "NULL" || v === "NEGATIVE"),
		getAnaliseRow("Кетоны(KET)", analise.ket, (v) => v === "NEG" || v === "NEGATIVE"),
		getAnaliseRow("Глюкоза(GLU)", analise.glu, (v) => v === "NEG" || v === "NEGATIVE"),
		getAnaliseRow("Кислотность(pH)", analise.ph, (v) => v >= 5 && v <= 7 || v === undefined),
		getAnaliseRow("Плотность(SG)", analise.sg, (v) => v >= 1.015 && v <= 1.04 || v === undefined),
		getAnaliseRow("Лейкоциты(LEU)", analise.leu),);

	return box;
};

const getAnaliseRow = (name, value, valid) => {
	const row = document.createElement("div");
	const nameElement = document.createElement("div");
	const valueElement = document.createElement("div");

	nameElement.className = "analise-name";
	nameElement.textContent = `${name}:`;

	valueElement.className = "analise-value";
	valueElement.textContent = value;

	if (valid) {
		const isValid = valid(value);
		if (isValid === undefined) {
			row.classList.add("undefined");
		} else if (isValid) {
			row.classList.add("valid");
		} else {
			row.classList.add("invalid");
		}
	} else {
		row.classList.add("valid");
	}

	row.append(nameElement, valueElement);

	return row;
};

const goHomePage = () => {
	location.href = "/"
}

const showError = (message) => {
	const containerErrors = document.querySelector('#container-notify')
	const containerElement = document.createElement("div")
	const headerElement = document.createElement("h3")
	const messageElement = document.createElement("p")

	containerElement.className = "error-container"
	headerElement.textContent = "😡 Произошла ошибка"
	messageElement.textContent = message


	containerElement.onclick = () => {
		containerElement.remove()
	}

	containerElement.append(
		headerElement,
		messageElement,
	)

	containerErrors.append(containerElement)
	setTimeout(() => containerElement.remove(), 2000)
}

const showSuccess = (message) => {
	const containerErrors = document.querySelector('#container-notify')
	const containerElement = document.createElement("div")
	const headerElement = document.createElement("h3")
	const messageElement = document.createElement("p")

	containerElement.className = "success-container"
	headerElement.textContent = "😊 Операция успешно выполнена"
	messageElement.textContent = message

	containerElement.onclick = () => {
		containerElement.remove()
	}

	containerElement.append(
		headerElement,
		messageElement,
	)

	containerErrors.append(containerElement)
	setTimeout(() => containerElement.remove(), 2000)
}

const onCreateUserClick = () => {
	openCreateUserModal()
}

const onCreateUser = () => {
	const nameElement = document.querySelector('#createUserName')
	const surnameElement = document.querySelector('#createUserSurname')
	const patronymicElement = document.querySelector('#createUserPatronymic')
	const snilsElement = document.querySelector('#createUserSnils')

	if (
		nameElement.value === "" ||
		surnameElement.value === "" ||
		patronymicElement.value === "" ||
		snilsElement.value === ""
	) {
		showError("Не все поля заполнены")
		return
	}

	fetch.createUser(
		nameElement.value,
		surnameElement.value,
		patronymicElement.value,
		snilsElement.value,
		() => {
			showSuccess("Пользователь успешно создан")
			"" === nameElement.textContent;
			closeCreateUserModal()
			loadUsers()
		},
	)
}

const onCloseCreateUser = () => {
	closeCreateUserModal()
}

const openCreateUserModal  = () => {
	const modal = document.querySelector(".create-user")
	modal.classList.remove("hidden")
}

const closeCreateUserModal  = () => {
	const modal = document.querySelector(".create-user")
	modal.classList.add("hidden")
}

const selectUser = document.getElementById('userSelect');
selectUser.classList.add('hidden');

selectUser.addEventListener('click', () => {
	selectUser.classList.toggle('visible');
});


loadUsers();
initAutoSearchUsers();



