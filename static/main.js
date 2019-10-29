let Suggests = {
    data() {
        return {
            open: false,
            current: 0,
            selected: '',
            name: this.sname,
            list:[],
        }
    },

    props: {
        sname: {
            type: String,
            required: true
        },

        lname: {
            type: String,
            required: true
        },

        suggestList: {
            type: Array,
            required: true,
            twoWay: true,
        },
    },

    computed: {
        matches() {
            return this.suggestList.filter((str) => {
                return str.indexOf(this.selected) >= 0;
            });
        },
        openSuggestion() {
            return this.selected !== "" &&
                this.matches.length !== 0 &&
                this.open === true;
        }
    },

    methods: {
        enter() {
            this.selected = this.matches[this.current];
            this.open = false;
        },
        up() {
            if(this.current > 0)
                this.current--;
        },
        down() {
            if(this.current < this.matches.length - 1)
                this.current++;
        },
        isActive(index) {
            return index === this.current;
        },
        emitChange(event) {
            let currentInput = event.target.value;
            if (currentInput.length === 3) {
                this.$emit("get-suggests", this.name, currentInput);
            }
            this.$emit('change-name', event.target.value);
        },
        change(event) {
            let currentInput = event.target.value;
            if (currentInput.length === 3) {
                this.$emit("get-suggests", this.name, currentInput);
            }

        },
        suggestionClick(index) {
            this.selected = this.matches[index];
            this.open = false;
            this.$emit('change-name', this.selected);
        },
        handleBlur() {
            this.open = false;
        },
    },

    template: `<div class="dropdown form-group" style="position:relative" v-bind:class="{'open':openSuggestion}">
    <label>{{lname}}</label>
    <input class="form-control dropdown-toggle" data-toggle="dropdown" type="text" v-model="selected"
        @keydown.enter = 'enter'
        @keydown.down = 'down'
        @keydown.up = 'up'
        @keydown.esc = "handleBlur"
        @input="emitChange"
        @change="emitChange"
        @blur="handleBlur"
    >
    <ul class="dropdown-menu" style="width:100%">
    <li v-for="(suggestion, index) in matches"
            v-bind:class="{'active': isActive(index)}"
            @click="suggestionClick(index)"
        ><a href="#">{{suggestion}}</a></li>
    </ul>
</div>`
};


new Vue({
    el: '#app',
    components: {suggests: Suggests},
    data: {
        builderName: '',
        modelName: '',
        buildersList: [],
        modelsList: [],
        boatsList: [],
        errorMsg: ''
    },

    methods: {
        getSuggestsParent: function(name, currentInput) {
            axios.get(`/v1/suggest?param=${name}&prefix=${currentInput}`)
                .then((response) => {
                    if (name === "builders") {
                        this.buildersList = response.data;
                    } else if (name === "models") {
                        this.modelsList = response.data;
                    }
                })
                .catch((error) => {
                    console.log(error);
                });
        },
        changeBuilderName: function(newName) {
            this.builderName = newName;
        },
        changeModelName: function(newName) {
            this.modelName = newName;
        },
        onFind: function(event, data) {
            if (event) event.preventDefault();
            axios.get(`/v1/boats/find?builder=${this.builderName}&model=${this.modelName}&limit=5`)
                .then((response) => {
                    if (response.data.yachts === null) {
                        this.errorMsg = "No yachts found";
                        return
                    }
                    this.errorMsg = "";
                    this.boatsList = response.data.yachts;
                })
                .catch((error) => {
                    console.log(error);
                });
        },
    },
});
