const errorElement = document.querySelector("#errorMessage");

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

	searchUsers = (search, callback) => {
		this.GET(`/users/${search}`, callback);
	};

	getAnalysis = (user, callback) => {
		this.GET(`/analyses/${user}`, callback);
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
			if (xhr.status === 403 ) {
				errorElement.textContent = 'Неверный пароль/логин';
				return;
			}
			if (xhr.response) {
				errorElement.textContent = '';
				callback(JSON.parse(xhr.response));
				return;
			}
			callback();
		};
		xhr.send(JSON.stringify(params));
	};
}
const fetch = new Fetch("http://localhost:8080");

function onSignIn() {
	const usernameElement = document.querySelector("#usernameInput");
	const passwordElement = document.querySelector("#passwordInput");
	const errorElement = document.querySelector("#errorMessage");

	const username = usernameElement.value.trim();
	const password = passwordElement.value.trim();

	if (username === '' || password === '' )
	{
		errorElement.textContent = 'Необходимо заполнить все поля!';
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

	box.append(
		getAnaliseRow("Дата и время сдачи", analise.date),
		getAnaliseRow("Эритроциты(Bld)", analise.bld),
		getAnaliseRow("Уробилиноген(UBG)", analise.ubg, (v) => v >= 0 && v <= 6),
		getAnaliseRow("Билирубин(Bil)", analise.bil, (v) => v >= 1 && v <= 3),
		getAnaliseRow("Белок(Pro)", analise.pro, (v) => v >= 0 && v <= 0.3),
		getAnaliseRow("Нитриты(Nit)", analise.nit),
		getAnaliseRow("Кетоны(KET)", analise.ket),
		getAnaliseRow("Глюкоза(GLU)", analise.glu),
		getAnaliseRow("Кислотность(pH)", analise.ph, (v) => v >= 5 && v <= 7),
		getAnaliseRow("Плотность(SG)", analise.sg, (v) => v >= 1.015 && v <= 1.04),
		getAnaliseRow("Лейкоциты(LEU)", analise.leu),
	);

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
		isValid ? row.classList.add("valid") : row.classList.add("invalid");
	} else {
		row.classList.add("valid");
	}

	row.append(nameElement, valueElement);

	return row;
};

function exit()
{
	const ex = document.querySelector("buttonExit")
location.href = "MainTemplate.html"


}
loadUsers();
initAutoSearchUsers();



