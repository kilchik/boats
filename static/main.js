var app = new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        builderName: '',
        modelName: '',
        boatsList: [],
    },

    methods: {
            onFind: function(event, data) {
            if (event) event.preventDefault();
            axios.get(`/v1/boats/find?builder=${this.builderName}&model=${this.modelName}&limit=5`)
                .then((response) => {
                    console.log('response', response);
                    this.boatsList = response.data.yachts;
                })
                .catch((error) => {
                    console.log(error);
                });
        },
    },
});
