const registration_form = {
    data() {
        return {
            address_same_checked: true,  
        }
    },
    props: ["items"],
    template: `
        <h3>Registration</h3>
        <hr>
        <form autocomplete="off" class="needs-validation" novalidate>
            <text_input label="First name" name="first-name" required="true" type="text"></text_input>
            <text_input label="Last name" name="last-name" required="true" type="text"></text_input>
            <text_input label="Password" name="password" required="true" type="password"></text_input>
            <text_input label="Email" name="email" required="true" type="email"></text_input>
            <select_input label="Favourite colour" name="Color" :items="items"></select_input>

            <text_input label="Address" name="address" required="true" type="text"></text_input>
            <text_input label="City/Town" name="city" required="true" type="text"></text_input>
            <text_input label="State/Province" name="province" required="true" type="text"></text_input>
            <text_input label="Zip Code" name="zip" required="true" type="text"></text_input>

            <check_input v-on:click="address_same" label="Mailing address same" checked="true" v-model="address_same_checked "</check_input>

            <div v-if="address_same_checked === false">
                <div class="mt-3">
                    <text_input label="Mailing Address" name="address2" type="text"></text_input>
                    <text_input label="Mailing City/Town" name="city2" type="text"></text_input>
                    <text_input label="Mailing State/Province" name="province2" type="text"></text_input>
                    <text_input label="Mailing  Zip Code" name="zip2" type="text"></text_input>
                </div>
            </div> 

            <check_input label="I agree to terms and conditions" required="true"</check_input>
            <hr>
            <input class="btn btn-outline-primary" type="submit" value="Register">
        </form>
    `,
    methods: {
        address_same() {
            console.log("Address same fired");
            if(this.address_same_checked === true) {
                console.log("Was checked on click.");
                this.address_same_checked = false;
            } else {
                console.log("Was not checked on click. ");
                this.address_same_checked = true; 
            }
        }
    },
    components: {
        'text_input': text_input,
        'select_input': select_input,
        'check_input': check_input, 
    },
    mounted () {
        (function () {
        'use strict'

        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        const forms = document.querySelectorAll('.needs-validation')

        // Loop over them and prevent submission
        Array.from(forms).forEach(form => {
            form.addEventListener('submit', event => {
            if (!form.checkValidity()) {
                event.preventDefault()
                event.stopPropagation()
            }

            form.classList.add('was-validated')
            }, false)
        })
        })()
    }
}