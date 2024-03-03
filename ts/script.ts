"use strict"

// let jsonData = JSON.parse('{"monday": {"first": [], "second": [{"sbj": "Проектная деятельность https://docs.google.com/spreadsheets/u/1/d/e/2PACX-1vSybcuU7Cv0_IGEg8sP7LD_mxQYu3akGUj_xxKX-5gXtdqcwAeDhtWRM8d4WGqscS3_LIQBWUThqoXk/pubhtml?gid=0&single=true", "teacher": "", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["*ПД*"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}, "tuesday": {"first": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Обучение LMS", "type": "Лекция"}], "fourth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["МСпорт. зал (графики кафедры)"], "location": "Михалковская ", "type": "Практика"}], "fifth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-10-31", "shortRooms": ["--МСпортзал--"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "seventh": []}, "wednesday": {"first": [], "second": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-09-01", "dt": "2023-10-15", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Навыки эффективной презентации", "teacher": "Олейникова Елизавета Витальевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Практика"}], "fourth": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 4, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [{"sbj": "Иностранный язык", "teacher": "Полякова Татьяна Владимировна, Чернякова Ирина Александровна, Черкасова Инна Петровна, Шляхтенков Юрий Григорьевич, Кожухова Валентина Валерьевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав5307", "ав5306", "ав5305", "ав5303", "ав5301"], "location": "Автозаводская", "type": "Практика"}], "sixth": [], "seventh": []}, "thursday": {"first": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "second": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-10-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-10-26", "dt": "2023-12-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}], "fourth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Васюткин Александр Олегович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-11-26", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-28", "dt": "2023-12-24", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}], "seventh": []}, "friday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-09-10", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}, {"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-11", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "third": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-25", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fourth": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fifth": [], "sixth": [], "seventh": []}, "saturday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр 2413 (ФО-2)"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Селюков Александр Сергеевич", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2305"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}}')
let jsonData = JSON.parse('{"monday": {"first": [], "second": [{"sbj": "Проектная деятельность https://docs.google.com/spreadsheets/u/1/d/e/2PACX-1vSybcuU7Cv0_IGEg8sP7LD_mxQYu3akGUj_xxKX-5gXtdqcwAeDhtWRM8d4WGqscS3_LIQBWUThqoXk/pubhtml?gid=0&single=true", "teacher": "", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["*ПД*"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}, "tuesday": {"first": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Муханов Сергей Александрович", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Webinar", "type": "Лекция"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-11-01", "shortRooms": [], "location": "Обучение LMS", "type": "Лекция"}], "fourth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["МСпорт. зал (графики кафедры)"], "location": "Михалковская ", "type": "Практика"}], "fifth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-10-31", "shortRooms": ["--МСпортзал--"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Вакансия_Инфобез 6", "df": "2023-11-05", "dt": "2023-11-19", "shortRooms": ["ав2205"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-11-20", "dt": "2023-12-24", "shortRooms": ["ав2216"], "location": "Автозаводская", "type": "Лаб. работа"}], "seventh": []}, "wednesday": {"first": [], "second": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна", "df": "2023-09-01", "dt": "2023-10-15", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Навыки эффективной презентации", "teacher": "Олейникова Елизавета Витальевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Практика"}], "fourth": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 4, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [{"sbj": "Иностранный язык", "teacher": "Полякова Татьяна Владимировна, Чернякова Ирина Александровна, Черкасова Инна Петровна, Шляхтенков Юрий Григорьевич, Кожухова Валентина Валерьевна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав5307", "ав5306", "ав5305", "ав5303", "ав5301"], "location": "Автозаводская", "type": "Практика"}], "sixth": [], "seventh": []}, "thursday": {"first": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "second": [{"sbj": "Безопасность операционных систем Windows", "teacher": "Вакансия_Инфобез 1, Морозов Алексей Константинович", "df": "2023-09-01", "dt": "2023-11-12", "shortRooms": ["ав4809"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-11-13", "dt": "2023-12-17", "shortRooms": ["ав4810"], "location": "Автозаводская", "type": "Лаб. работа"}], "third": [{"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-09-01", "dt": "2023-10-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}, {"sbj": "Введение в аналитику информационной безопасности", "teacher": "Кривоногов Антон Алексеевич, Плоткин Александр Сергеевич", "df": "2023-10-26", "dt": "2023-12-24", "shortRooms": ["ав3105"], "location": "Автозаводская", "type": "Лаб. работа"}], "fourth": [{"sbj": "Методы и средства криптографической защиты информации", "teacher": "Бутакова Наталья Георгиевна, Васюткин Александр Олегович", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["ав4805"], "location": "Автозаводская", "type": "Лаб. работа"}], "fifth": [], "sixth": [{"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-09-01", "dt": "2023-11-26", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}, {"sbj": "Общая физическая подготовка (см. график кафедры)", "teacher": "", "df": "2023-11-28", "dt": "2023-12-24", "shortRooms": ["М спортзал"], "location": "Михалковская ", "type": "Практика"}], "seventh": []}, "friday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-09-10", "shortRooms": ["Пр2503"], "location": "Прянишникова", "type": "Лаб. работа"}, {"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-11", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "third": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-09-25", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fourth": [{"sbj": "Сети и системы передачи информации", "teacher": "Карпов Александр Викторович, Камозин Сергей Андреевич", "df": "2023-11-02", "dt": "2023-12-24", "shortRooms": ["Пр2402"], "location": "Прянишникова", "type": "Лаб. работа"}], "fifth": [], "sixth": [], "seventh": []}, "saturday": {"first": [{"sbj": "Основы веб-технологий", "teacher": "Гнибеда Артём Юрьевич, Энгиноева Диана Хизировна", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр 2413 (ФО-2)"], "location": "Прянишникова", "type": "Лаб. работа"}], "second": [{"sbj": "Линейная алгебра и функции нескольких переменных", "teacher": "Селюков Александр Сергеевич", "df": "2023-09-01", "dt": "2023-12-24", "shortRooms": ["Пр2305"], "location": "Прянишникова", "type": "Практика"}], "third": [], "fourth": [], "fifth": [], "sixth": [], "seventh": []}}')

// From (2023-09-01) to (01 Сен)
function convertDate(date: string): string {
    const monthNames = ["Янв", "Фев", "Мар", "Апр", "Мая", "Июн",
        "Июл", "Авг", "Сен", "Окт", "Ноя", "Дек"]
    let splittedDate = date.split("-")
    return splittedDate[2] + " " + monthNames[parseInt(splittedDate[1]) - 1]
}

function pairNumberToPairTime(pairNumber: string): string {
    let pairs: { [key: string]: string } = {
        "first": "09:00 - 10:30",
        "second": "10:40 - 12:10",
        "third": "12:20 - 13:50",
        "fourth": "14:30 - 16:00",
        "fifth": "16:10 - 17:40",
        "sixth": "17:50 - 19:20",
        "seventh": "19:30 - 21:00",
    }
    return pairs[pairNumber]
}

function addPair(pairData: { [key: string]: any }, pairNumber: string): HTMLDivElement {
    let pair = document.createElement("div")
    pair.setAttribute('class', `pair ${pairNumber}`)
    pair.innerHTML = `
                        <div class="date-time">
                            <div class="time">${pairNumberToPairTime(pairNumber)}</div>
                            <div class="date">${convertDate(pairData["df"])} - ${convertDate(pairData["dt"])}</div>
                        </div>
                        <div class="rooms">
                            <div class="room">${pairData["shortRooms"].map((room: string) => room.charAt(0).toUpperCase() + room.slice(1)).join(" ")}</div>
                        </div>
                        <div class="discipline">${pairData["sbj"]} (${pairData["type"]})</div>
                        <div class="teachers">
                            <div class="teacher">${pairData["teacher"].replace(", ", "<br>")}</div>
                        </div>
                     `

    return pair
}

function addTimeSlot(slotData: [{ [key: string]: any }], pairNumber: string): HTMLDivElement[] {
    let pairs: HTMLDivElement[] = []
    for (const data of slotData) {
        pairs.push(addPair(data, pairNumber))
    }

    return pairs
}

function addDay(dayData: { [key: string]: any }, dayName: string): HTMLDivElement[] {
    let pairs: HTMLDivElement[] = [];
    for (const pairNumber in dayData) {
        const timeSlots = addTimeSlot(dayData[pairNumber], pairNumber);
        pairs.push(...timeSlots);
    }
    return pairs
}

function addSchedule(jsonData: { [key: string]: object }): void {
    for (const data in jsonData) {
        const day = addDay(jsonData[data], data)
        for (const pair of day) {
            document.getElementById(data)?.appendChild(pair);
        }
    }
}

function closeSvgHandler() {
    const groupInput = document.querySelector(".group-input") as HTMLInputElement
    groupInput.value = ""

}

function burgerMenuHandler() {
    let target = document.querySelector('.burger-menu');
    if (!target) return
    if (target.classList.contains("closed")) {
        target.classList.remove("closed")
        target.classList.add("shown")
        const firstChild = target.children[0] as HTMLElement;
        const secondChild = target.children[1] as HTMLElement;
        const thirdChild = target.children[2] as HTMLElement;
        firstChild.style.transform = 'translateY(6px) rotate(45deg)'
        secondChild.style.opacity = '0'
        thirdChild.style.transform = 'translateY(-8px) rotate(-45deg)'
    }
    else {
        target.classList.remove("shown")
        target.classList.add("closed")
        const firstChild = target.children[0] as HTMLElement;
        const secondChild = target.children[1] as HTMLElement;
        const thirdChild = target.children[2] as HTMLElement;
        firstChild.style.transform = 'translateY(0px) rotate(0deg)'
        secondChild.style.opacity = '1'
        thirdChild.style.transform = 'translateY(0px) rotate(0deg)'
    }
}

function setDate(date: string) {
    const dateRegexp = /[0-9]{4}-[0-9]{2}-[0-9]{2}/
    const dateInput = document.querySelector(".date-input") as HTMLInputElement

    if (dateRegexp.test(date)) dateInput.value = date
    else {
        const today = new Date();
        const yyyy = today.getFullYear();
        const mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
        const dd = String(today.getDate()).padStart(2, '0');
        dateInput.value = yyyy + "-" + mm + "-" + dd
    }
    dateInputHandler()
}

function dateInputHandler() {
    const dateRegexp = /[0-9]{4}-[0-9]{2}-[0-9]{2}/
    const dateInput = document.querySelector(".date-input") as HTMLInputElement

    const today = new Date();
    const yyyy = today.getFullYear();
    const mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
    const dd = String(today.getDate()).padStart(2, '0');

    if (dateRegexp.test(dateInput.value) && dateInput.value >= yyyy + "-" + mm + "-" + dd) {
        const selectedDate = new Date(dateInput.value)
        const dates = document.getElementsByClassName("month-date")
        let dayNum = 0
        for (const date of dates) {
            let dayDate = selectedDate
            dayDate.setDate(selectedDate.getDate() + 1 + dayNum - (selectedDate.getDay() == 0 ? 6 : selectedDate.getDay()))
            date.textContent = String(dayDate.getDate())
            dayNum++
        }
    }
    else {
        setDate("")     
    }
}

window.onload = function () {
    addSchedule(jsonData)
    setDate("")

    const preloaderContainer = document.querySelector(".main-preloader-container") as HTMLElement

    setTimeout(() => { preloaderContainer.style.animation = "fadeOut 1s cubic-bezier(0.645, 0.045, 0.355, 1) 0.5s forwards"; }, 1);
    preloaderContainer.addEventListener('animationend', function () {
        preloaderContainer.style.display = 'none'
    })

    const closeSvg = document.querySelector(".close-svg") as HTMLInputElement
    if (closeSvg) closeSvg.onclick = closeSvgHandler

    const burgerMenu = document.querySelector(".burger-menu") as HTMLInputElement
    if (burgerMenu) burgerMenu.onclick = burgerMenuHandler

    const burgerMenuClose = document.querySelector(".burger-menu-close") as HTMLInputElement
    if (burgerMenuClose) burgerMenuClose.onclick = burgerMenuHandler

    const dateInput = document.querySelector(".date-input") as HTMLInputElement
    if (dateInput) dateInput.onchange = dateInputHandler
}
