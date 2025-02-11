<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-3">Login</h1>
                <hr>
                <form-tag>
                    
                    <text-input 
                        label="Email"
                        type="email"
                        name="email"
                        required="true">
                    </text-input>

                    <text-input 
                        label="Password"
                        type="password"
                        name="password"
                        required="true">
                    </text-input>  

                    <hr>
                    <input class="btn btn-primary" type="submit" value="Login">
                </form-tag>
            </div>
        </div>
    </div>
</template>

<script>
import TextInput from './forms/TextInput.vue'
import FormTag from './forms/FormTag.vue'

export default {
    name: 'AppLogin',
    components: { 
        TextInput,
        FormTag,
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
</script>