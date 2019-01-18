package webauthnkit.example

import android.content.Intent
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import org.jetbrains.anko.*
import org.jetbrains.anko.sdk27.coroutines.onClick

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        verticalLayout {

            padding = dip(10)

            button("Registration") {
                textSize = 24f

                onClick {
                    goToRegistrationActivity()
                }

            }

            button("Authentication") {
                textSize = 24f

                onClick {
                    goToAuthenticationActivity()
                }
            }

        }
    }

    private fun goToRegistrationActivity() {
        var intent = Intent(this, RegistrationActivity::class.java)
        startActivity(intent)
    }

    private fun goToAuthenticationActivity() {
        var intent = Intent(this, AuthenticationActivity::class.java)
        startActivity(intent)
    }
}
