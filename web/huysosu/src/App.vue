<script>
import { ref } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const selectedOption = ref('');
    const inputVisible = ref(false);
    const searchQuery = ref('');
    const searchResults = ref([]);
    const selectedStudent = ref(null);

    function showInput() {
      inputVisible.value = selectedOption.value !== '';
    }

    async function fetchDataFromAPI() {
      try {
        const apiUrl = 'http://localhost:8000/api/search';

        const axiosInstance = axios.create({
          baseURL: 'http://localhost:8080', // Замените на ваш URL
          headers: {
            'Content-Type': 'application/json; charset=utf-8',
          },
        });

        const response = await axiosInstance.post(apiUrl, {
          option: selectedOption.value,
          query: searchQuery.value,
        });
        searchResults.value = response.data;
        console.log(response.data)
      } catch (error) {
        console.error('Error fetching data from API:', error);
      }
    }

    function search() {
      fetchDataFromAPI();
    }

    function selectStudent(student) {
      selectedStudent.value = student;
    }

    return {
      selectedOption,
      inputVisible,
      showInput,
      searchQuery,
      searchResults,
      selectedStudent,
      search,
      selectStudent,
    };
  },
};
</script>

<template>
 <div id="BackSVG"> <svg width="1000" height="800" viewBox="0 0 1325 1004" fill="none" xmlns="http://www.w3.org/2000/svg">
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -235.683 -31.4908)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -211.116 116.292)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -95.5778 20.1185)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -186.55 264.076)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -71.0115 167.902)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 44.5272 71.7279)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 160.066 -24.4459)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -161.984 411.859)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -46.4455 315.685)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 69.0934 219.511)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 184.632 123.337)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 300.171 27.1635)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -252.956 655.816)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -137.418 559.642)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -21.8788 463.468)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 93.6598 367.294)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 209.199 271.12)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 324.737 174.947)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 440.276 78.7728)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 555.815 -17.401)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -228.39 803.599)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -112.851 707.425)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 2.68747 611.251)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 118.226 515.077)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 233.765 418.904)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 349.304 322.73)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 464.843 226.556)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 580.381 130.382)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 695.92 34.2084)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -203.823 951.382)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 -88.285 855.208)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 27.2538 759.034)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 142.793 662.861)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 258.331 566.687)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 373.87 470.513)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 489.409 374.339)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 604.948 278.165)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 720.486 181.992)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 836.025 85.8177)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  <rect x="0.853469" y="-0.147051" width="150.847" height="149.816" transform="matrix(0.768576 -0.639758 0.938361 0.345656 951.564 -10.356)" stroke="#0E0E0E" stroke-opacity="0.25"/>
  </svg></div>

  <header class="Head">
        <div id="flex-auth">
          <a class="butn" href="#">Выйти</a>
        </div>
      </header>
      <body class="Main">
        <div class="maindv">
          <span id="Students">Студенты</span>
          <div id="flex-tool">
            <button id="Search" @click="search">Поиск</button>
            <select id="MainList" v-model="selectedOption" @change="showInput">
              <option value="">Выбрать</option>
              <option value="full_name">ФИО</option>
              <option value="Номер карточки">Номер карточки</option>
              <option value="Дата Рождения">Дата Рождения</option>
              <option value="Дата зачисления">Дата зачисления</option>
              <option value="Адрес проживания">Адрес проживания</option>
              <option value="Место рождения">Место рождения</option>
              <option value="Номер приказа о предоставлении общежития">Номер приказа о предоставлении общежития</option>
              <option value="Номер приказа о зачислении">Номер приказа о зачислении</option>
            </select>
          </div>

          <form id="inputContainer" v-show="inputVisible" action="" method="get">
            <input v-model="searchQuery" id="into-value" name="name" type="text">
          </form>

          <div id="flex-boxes">
            <div class="SearchResults">
              <table id="table">
                <tr>
                  <th id="MainTr"> Список студентов: </th>
                </tr>
                <tbody>
                  <tr v-for="result in searchResults" :key="result.id" @click="selectStudent(result)">
                    <td>{{ result.full_name }} </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="MemberCard" v-if="selectedStudent">
              <h2>{{ selectedStudent.full_name }}</h2>
                <p>Номер карточки: {{ selectedStudent.card_number }}</p>
            </div>
          </div>
        </div>
      </body>
</template>
