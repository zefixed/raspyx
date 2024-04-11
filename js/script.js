"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
// let jsonData = JSON.parse('{"monday": {"first": [], "second": [{"sbj": "Проектная деятельность https://docs.google.com/spreadsheets/u/1/d/e/2PACX-1vSybcuU7Cv0_IGEg8sP7LD_mxQYu3akGUj_xxKX-5gXtdqcwAeDhtWRM8d4WGqscS3_LIQBWUThqoXk/pubhtml?gid=0&single=true", "teacher": "", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["*ПД*"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}, "tuesday": {"first": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Обучение LMS", "type": "Лекция"}], "fourth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["МСпорт. зал (графики кафедры)"], "location": "Михалковская ", "type": "Практика"}], "fifth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-10-31", "shortRooms": ["--МСпортзал--"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "seventh": []}, "wednesday": {"first": [], "second": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-09-01", "dt": "2023-10-15", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Навыки эффективной презентации", "teacher": "Олейникова Елизавета Витальевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Практика"}], "fourth": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 4, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [{"sbj": "Иностранный язык", "teacher": "Полякова Татьяна Владимировна, Чернякова Ирина Александровна, Черкасова Инна Петровна, Шляхтенков Юрий Григорьевич, Кожухова Валентина Валерьевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав5307", "ав5306", "ав5305", "ав5303", "ав5301"], "location": "Автозаводская", "type": "Практика"}], "sixth": [], "seventh": []}, "thursday": {"first": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "second": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-10-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-10-26", "dt": "2023-12-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}], "fourth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Васюткин Александр Олегович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-11-26", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-28", "dt": "2023-12-24", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}], "seventh": []}, "friday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-09-10", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}, {"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-11", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "third": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-25", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fourth": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fifth": [], "sixth": [], "seventh": []}, "saturday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр 2413 (ФО-2)"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Селюков Александр Сергеевич", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2305"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}}')
let jsonData = JSON.parse('{"monday": {"first": [], "second": [{"sbj": "Проектная деятельность https://docs.google.com/spreadsheets/u/1/d/e/2PACX-1vSybcuU7Cv0_IGEg8sP7LD_mxQYu3akGUj_xxKX-5gXtdqcwAeDhtWRM8d4WGqscS3_LIQBWUThqoXk/pubhtml?gid=0&single=true", "teacher": "", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["*ПД*"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}, "tuesday": {"first": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Обучение LMS", "type": "Лекция"}], "fourth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["МСпорт. зал (графики кафедры)"], "location": "Михалковская ", "type": "Практика"}], "fifth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-10-31", "shortRooms": ["--МСпортзал--"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "seventh": []}, "wednesday": {"first": [], "second": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-09-01", "dt": "2023-10-15", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Навыки эффективной презентации", "teacher": "Олейникова Елизавета Витальевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Практика"}], "fourth": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 4, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [{"sbj": "Иностранный язык", "teacher": "Полякова Татьяна Владимировна, Чернякова Ирина Александровна, Черкасова Инна Петровна, Шляхтенков Юрий Григорьевич, Кожухова Валентина Валерьевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав5307", "ав5306", "ав5305", "ав5303", "ав5301"], "location": "Автозаводская", "type": "Практика"}], "sixth": [], "seventh": []}, "thursday": {"first": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "second": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-10-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-10-26", "dt": "2023-12-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}], "fourth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Васюткин Александр Олегович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-11-26", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-28", "dt": "2023-12-24", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}], "seventh": []}, "friday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-09-10", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}, {"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-11", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "third": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-25", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fourth": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fifth": [], "sixth": [], "seventh": []}, "saturday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр 2413 (ФО-2)"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Селюков Александр Сергеевич", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2305"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}}');
let groups;
(() => __awaiter(void 0, void 0, void 0, function* () { groups = (yield getGroups()).groups; }))();
// From (2023-09-01) to (01 Сен)
function convertDate(date) {
    const monthNames = ["Янв", "Фев", "Мар", "Апр", "Мая", "Июн",
        "Июл", "Авг", "Сен", "Окт", "Ноя", "Дек"];
    let splittedDate = date.split("-");
    return splittedDate[2] + " " + monthNames[parseInt(splittedDate[1]) - 1];
}
function pairNumberToPairTime(pairNumber) {
    let pairs = {
        "first": "09:00 - 10:30",
        "second": "10:40 - 12:10",
        "third": "12:20 - 13:50",
        "fourth": "14:30 - 16:00",
        "fifth": "16:10 - 17:40",
        "sixth": "17:50 - 19:20",
        "seventh": "19:30 - 21:00",
    };
    return pairs[pairNumber];
}
function addPair(pairData, pairNumber) {
    let pair = document.createElement("div");
    pair.setAttribute('class', `pair ${pairNumber}`);
    pair.innerHTML = `
                        <div class="date-time">
                            <div class="time">${pairNumberToPairTime(pairNumber)}</div>
                            <div class="date">${convertDate(pairData["df"])} - ${convertDate(pairData["dt"])}</div>
                        </div>
                        <div class="rooms">
                            <div class="room">${pairData["shortRooms"].map((room) => room.charAt(0).toUpperCase() + room.slice(1)).join(" ")}</div>
                        </div>
                        <div class="discipline">${pairData["sbj"]} (${pairData["type"]})</div>
                        <div class="teachers">
                            <div class="teacher">${pairData["teacher"].replace(", ", "<br>")}</div>
                        </div>
                     `;
    return pair;
}
function addTimeSlot(slotData, pairNumber) {
    let pairs = [];
    for (const data of slotData) {
        pairs.push(addPair(data, pairNumber));
    }
    return pairs;
}
function addDay(dayData, dayName) {
    let pairs = [];
    for (const pairNumber in dayData) {
        const timeSlots = addTimeSlot(dayData[pairNumber], pairNumber);
        pairs.push(...timeSlots);
    }
    return pairs;
}
function addSchedule(jsonData) {
    var _a;
    for (const data in jsonData) {
        const day = addDay(jsonData[data], data);
        for (const pair of day) {
            (_a = document.getElementById(data)) === null || _a === void 0 ? void 0 : _a.appendChild(pair);
        }
    }
}
function closeSvgHandler() {
    const groupInput = document.querySelector(".group-input");
    groupInput.value = "";
}
function burgerMenuHandler() {
    let target = document.querySelector('.burger-menu');
    if (!target)
        return;
    if (target.classList.contains("closed")) {
        target.classList.remove("closed");
        target.classList.add("shown");
        const firstChild = target.children[0];
        const secondChild = target.children[1];
        const thirdChild = target.children[2];
        firstChild.style.transform = 'translateY(6px) rotate(45deg)';
        secondChild.style.opacity = '0';
        thirdChild.style.transform = 'translateY(-8px) rotate(-45deg)';
    }
    else {
        target.classList.remove("shown");
        target.classList.add("closed");
        const firstChild = target.children[0];
        const secondChild = target.children[1];
        const thirdChild = target.children[2];
        firstChild.style.transform = 'translateY(0px) rotate(0deg)';
        secondChild.style.opacity = '1';
        thirdChild.style.transform = 'translateY(0px) rotate(0deg)';
    }
}
function setDate(newDate) {
    const datePattern = /^\d{4}-\d{2}-\d{2}$/;
    const dateInput = document.querySelector(".date-input");
    const currentDate = new Date();
    const formattedDate = `${currentDate.getFullYear()}-${String(currentDate.getMonth() + 1).padStart(2, '0')}-${String(currentDate.getDate()).padStart(2, '0')}`;
    dateInput.value = datePattern.test(newDate) ? newDate : formattedDate;
    dateInputHandler();
}
function dateInputHandler() {
    const dateRegexp = /[0-9]{4}-[0-9]{2}-[0-9]{2}/;
    const dateInput = document.querySelector(".date-input");
    const today = new Date();
    const yyyy = today.getFullYear();
    const mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
    const dd = String(today.getDate()).padStart(2, '0');
    if (dateRegexp.test(dateInput.value) && dateInput.value >= yyyy + "-" + mm + "-" + dd) {
        const selectedDate = new Date(dateInput.value);
        const dates = document.getElementsByClassName("month-date");
        let dayNum = 0;
        for (const date of dates) {
            let dayDate = selectedDate;
            dayDate.setDate(selectedDate.getDate() + 1 + dayNum - (selectedDate.getDay() == 0 ? 7 : selectedDate.getDay()));
            date.textContent = String(dayDate.getDate());
            dayNum++;
        }
    }
    else {
        setDate("");
    }
}
function setVersion(version) {
    const settingsModalFooter = document.querySelector(".settingsModalFooter");
    settingsModalFooter.textContent = `v${version}`;
}
function switchDarkThemeHandler() {
    var _a, _b;
    const switchTheme = document.querySelector("#flexSwitchCheckDarkTheme");
    // Set theme
    let theme = "light";
    if (switchTheme.checked)
        theme = "dark";
    const nodes = {
        body: document.querySelector(".body"),
        dateInput: document.querySelector(".date-input"),
        groupInput: document.querySelector(".group-input"),
        peopleSvg: document.querySelector(".people-svg"),
        closeSvg: document.querySelector(".close-svg"),
        days: document.querySelector(".days"),
        pairs: document.querySelector(".pairs"),
        day: document.getElementsByClassName("day"),
        pair: document.getElementsByClassName("pair"),
        modalContent: document.querySelector(".modal-content"),
        autocomplete: document.getElementsByClassName("autocomplete"),
        autocompleteItem: document.getElementsByClassName("autocomplete-item"),
    };
    // Remove/set dark classes
    if (theme == "light") {
        nodes["body"].classList.remove("body-dark");
        nodes["dateInput"].classList.remove("date-input-dark");
        nodes["groupInput"].classList.remove("group-input-dark");
        nodes["peopleSvg"].classList.remove("people-svg-dark");
        nodes["closeSvg"].classList.remove("close-svg-dark");
        nodes["days"].classList.remove("days-dark");
        nodes["pairs"].classList.remove("pairs-dark");
        for (const node of nodes["day"])
            node.classList.remove("day-dark");
        for (const node of nodes["pair"])
            node.classList.remove("pair-dark");
        nodes["modalContent"].classList.remove("modal-content-dark");
        for (const node of nodes["autocomplete"])
            node.classList.remove("autocomplete-dark");
        for (const node of nodes["autocompleteItem"])
            node.classList.remove("autocomplete-item-dark");
        document.body.style.setProperty('--light-selection-bg-color', "#6E6E6E");
        document.body.style.setProperty('--light-add-font-color', "#fffefe");
    }
    else {
        nodes["body"].classList.add("body-dark");
        nodes["dateInput"].classList.add("date-input-dark");
        nodes["groupInput"].classList.add("group-input-dark");
        nodes["peopleSvg"].classList.add("people-svg-dark");
        nodes["closeSvg"].classList.add("close-svg-dark");
        nodes["days"].classList.add("days-dark");
        nodes["pairs"].classList.add("pairs-dark");
        for (const node of nodes["day"])
            node.classList.add("day-dark");
        for (const node of nodes["pair"])
            node.classList.add("pair-dark");
        nodes["modalContent"].classList.add("modal-content-dark");
        for (const node of nodes["autocomplete"])
            node.classList.add("autocomplete-dark");
        for (const node of nodes["autocompleteItem"])
            node.classList.add("autocomplete-item-dark");
        document.body.style.setProperty('--light-selection-bg-color', "#fffefe");
        document.body.style.setProperty('--light-add-font-color', "#6E6E6E");
    }
    // Set data-bs-theme
    const html = document.getElementsByTagName("html").item(0);
    html === null || html === void 0 ? void 0 : html.setAttribute("data-bs-theme", theme);
    // Set icons
    const peopleSvg = (_a = document.querySelector(".people-svg")) === null || _a === void 0 ? void 0 : _a.children[0];
    const closeSvg = (_b = document.querySelector(".close-svg")) === null || _b === void 0 ? void 0 : _b.children[0];
    if (theme == "light") {
        peopleSvg === null || peopleSvg === void 0 ? void 0 : peopleSvg.setAttribute("src", "images/on_white/people_blue.svg");
        closeSvg === null || closeSvg === void 0 ? void 0 : closeSvg.setAttribute("src", "images/on_white/close.svg");
    }
    else {
        peopleSvg === null || peopleSvg === void 0 ? void 0 : peopleSvg.setAttribute("src", "images/on_black/bluePeople.svg");
        closeSvg === null || closeSvg === void 0 ? void 0 : closeSvg.setAttribute("src", "images/on_black/Exit.svg");
    }
}
function switchFooterHandler() {
    var _a, _b;
    const footer = document.querySelector(".footer");
    const switchFooter = document.querySelector("#flexSwitchCheckFooter");
    const pairs = document.querySelector(".pairs");
    if (switchFooter.checked) {
        footer.style.display = "none";
        for (const day of (_a = pairs === null || pairs === void 0 ? void 0 : pairs.children) !== null && _a !== void 0 ? _a : []) {
            day.classList.remove("day-with-footer");
            day.classList.add("day-without-footer");
        }
    }
    else {
        footer.style.display = "grid";
        for (const day of (_b = pairs === null || pairs === void 0 ? void 0 : pairs.children) !== null && _b !== void 0 ? _b : []) {
            day.classList.remove("day-without-footer");
            day.classList.add("day-with-footer");
        }
    }
}
function getGroups() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const response = yield fetch("https://itcpd.ru/api/V1/getGroups/?format=json");
            const data = yield response.json();
            return data;
        }
        catch (error) {
            console.log(error);
        }
    });
}
function setGroupsAutocomplete(groupName = "") {
    return __awaiter(this, void 0, void 0, function* () {
        const groupAutocompleteItems = document.querySelector(".autocomplete-items");
        const groupAutocomplete = document.querySelector(".group-autocomplete");
        const switchTheme = document.querySelector("#flexSwitchCheckDarkTheme");
        let theme = "light";
        if (switchTheme.checked)
            theme = "dark";
        groupAutocompleteItems.innerHTML = "";
        groupAutocomplete.style.width = `calc(100% + 1px)`; // idfn where that 1px comes from
        for (const group in groups) {
            if (group.split(/\-|\s/).join("").toLowerCase().includes(groupName.split(/\-|\s/).join("").toLowerCase())) {
                const listItem = document.createElement("li");
                listItem.classList.add("autocomplete-item");
                if (theme == "dark")
                    listItem.classList.add("autocomplete-item-dark");
                listItem.textContent = group;
                groupAutocompleteItems.appendChild(listItem);
            }
        }
    });
}
function groupAutocompleteHandler(event) {
    const groupAutocomplete = document.querySelector(".group-autocomplete");
    const groupInput = document.querySelector(".group-input");
    groupAutocomplete.style.display = "block";
    setGroupsAutocomplete(groupInput.value);
}
function setGroupHandler(event) {
    var _a;
    const groupInput = document.querySelector(".group-input");
    const groupAutocomplete = document.querySelector(".group-autocomplete");
    const target = event.target;
    if (target.tagName === "LI") {
        groupInput.value = (_a = target.textContent) !== null && _a !== void 0 ? _a : "";
        groupAutocomplete.style.display = "none";
        setSchedule(groupInput.value);
    }
    else {
        console.log("Easter egg for @sleepqeelz");
    }
}
function bodyHandler(event) {
    const list = Array.from(document.getElementsByTagName("ul"));
    const target = event.target;
    if (target.tagName == "INPUT")
        return;
    for (const element of list) {
        element.style.display = "none";
    }
}
function setScheduleHandler(event) {
    if (event.key == "Enter") {
        const groupInput = document.querySelector(".group-input");
        for (const group in groups) {
            if (group.split(/\-|\s/).join("").toLowerCase().includes(groupInput.value.split(/\-|\s/).join("").toLowerCase())) {
                groupInput.value = group;
                for (const element of Array.from(document.getElementsByTagName("ul"))) {
                    element.style.display = "none";
                }
                setSchedule(group);
                break;
            }
        }
    }
}
function setSchedule(group) {
    const pairs = document.querySelector(".pairs");
    const switchTheme = document.querySelector("#flexSwitchCheckDarkTheme");
    let theme = "light";
    if (switchTheme.checked)
        theme = "dark";
    const noRasp = `<div class="pair-none d-flex justify-content-center align-items-center">
            Сегодня нет занятий
        </div>`;
    pairs.innerHTML =
        `<div id="monday" class="day day-with-footer p-monday m-left mn-wdth">${noRasp}</div>
        <div id="tuesday" class="day day-with-footer p-tuesday m-all mn-wdth">${noRasp}</div>
        <div id="wednesday" class="day day-with-footer p-wednesday m-all mn-wdth">${noRasp}</div>
        <div id="thursday" class="day day-with-footer p-thursday m-all mn-wdth">${noRasp}</div>
        <div id="friday" class="day day-with-footer p-friday m-all mn-wdth">${noRasp}</div>
        <div id="saturday" class="day day-with-footer p-saturday m-right mn-wdth">${noRasp}</div>`;
    if (theme == "dark") {
        for (const day of document.getElementsByClassName("day")) {
            day.classList.add("day-dark");
        }
    }
}
window.onload = function () {
    setGroupsAutocomplete();
    addSchedule(jsonData);
    setDate("");
    switchDarkThemeHandler();
    setVersion("1.24.15");
    // Clearing the group's input 
    const groupInput = document.querySelector(".group-input");
    if (groupInput)
        groupInput.value = "";
    // Preloader
    const preloaderContainer = document.querySelector(".main-preloader-container");
    setTimeout(() => { preloaderContainer.style.animation = "fadeOut 1s cubic-bezier(0.645, 0.045, 0.355, 1) 0.5s forwards"; }, 1);
    preloaderContainer.addEventListener('animationend', function () {
        preloaderContainer.style.display = 'none';
    });
    // Group input clean
    const closeSvg = document.querySelector(".close-svg");
    if (closeSvg)
        closeSvg.onclick = closeSvgHandler;
    // Burger menu open
    const burgerMenu = document.querySelector(".burger-menu");
    if (burgerMenu)
        burgerMenu.onclick = burgerMenuHandler;
    // Burger menu close
    const burgerMenuClose = document.querySelector(".burger-menu-close");
    if (burgerMenuClose)
        burgerMenuClose.onclick = burgerMenuHandler;
    // Date input
    const dateInput = document.querySelector(".date-input");
    if (dateInput)
        dateInput.onchange = dateInputHandler;
    // Dark theme switch
    const switchDarkTheme = document.querySelector("#flexSwitchCheckDarkTheme");
    if (switchDarkTheme)
        switchDarkTheme.onclick = switchDarkThemeHandler;
    // Show footer switch
    const switchFooter = document.querySelector("#flexSwitchCheckFooter");
    if (switchFooter)
        switchFooter.onclick = switchFooterHandler;
    // Show group autocomplete
    const groupAutocomplete = document.querySelector(".group-input-group");
    if (groupAutocomplete)
        groupAutocomplete.onclick = groupAutocompleteHandler;
    if (groupAutocomplete)
        groupAutocomplete.onfocus = groupAutocompleteHandler;
    if (groupAutocomplete)
        groupAutocomplete.oninput = groupAutocompleteHandler;
    if (groupAutocomplete)
        groupAutocomplete.onkeyup = setScheduleHandler;
    const autocompleteItems = document.querySelector(".autocomplete-items");
    if (autocompleteItems)
        autocompleteItems.onclick = setGroupHandler;
    // Closing all drop-down lists
    const body = document.querySelector(".body");
    if (body)
        body.onclick = bodyHandler;
};
