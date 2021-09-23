package webauthnkit.core.authenticator.internal.ui

import android.annotation.TargetApi
import android.app.Activity.RESULT_OK
import android.app.KeyguardManager
import android.content.Context.KEYGUARD_SERVICE
import android.content.Intent
import android.os.Build
import android.text.format.DateFormat
import androidx.fragment.app.FragmentActivity

import kotlin.coroutines.resume
import kotlin.coroutines.suspendCoroutine
import kotlin.coroutines.Continuation
import kotlin.coroutines.resumeWithException

import kotlinx.coroutines.ExperimentalCoroutinesApi

import webauthnkit.core.authenticator.internal.PublicKeyCredentialSource
import webauthnkit.core.authenticator.internal.ui.dialog.*
import webauthnkit.core.error.*
import webauthnkit.core.util.WAKLogger
import webauthnkit.core.data.*
import java.util.*

@ExperimentalUnsignedTypes
@ExperimentalCoroutinesApi
@TargetApi(Build.VERSION_CODES.M)
class DefaultUserConsentUI(
    private val activity: FragmentActivity
) : UserConsentUI {

    companion object {
        val TAG = DefaultUserConsentUI::class.simpleName
        const val REQUEST_CODE = 6749
    }

    var keyguardResultListener: KeyguardResultListener? = null

    override val config = UserConsentUIConfig()

    override var isOpen: Boolean = false
        private set

    private var cancelled: ErrorReason? = null

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?): Boolean {

        WAKLogger.d(TAG, "onActivityResult")

        return if (requestCode == REQUEST_CODE) {

            WAKLogger.d(TAG, "This is my result")

            keyguardResultListener?.let {
                if (resultCode == RESULT_OK) {
                    WAKLogger.d(TAG, "OK")
                    it.onAuthenticated()
                } else {
                    WAKLogger.d(TAG, "Failed")
                    it.onFailed()
                }
            }
            keyguardResultListener = null
            true

        } else {
            false
        }
    }

    private fun <T> finish(cont: Continuation<T>, result: T) {
        cont.resume(result)
    }

    override fun cancel(reason: ErrorReason) {
        cancelled = reason
    }

    override suspend fun requestUserConsent(
        rpEntity: PublicKeyCredentialRpEntity,
        userEntity: PublicKeyCredentialUserEntity,
        requireUserVerification: Boolean
    ): String = suspendCoroutine { cont ->
        finish(cont, getDefaultKeyName(userEntity.name))
    }

    override suspend fun requestUserSelection(
        sources: List<PublicKeyCredentialSource>,
        requireUserVerification: Boolean
    ): PublicKeyCredentialSource = suspendCoroutine { cont ->
        finish(cont, sources.first())
    }

    private fun getDefaultKeyName(username: String): String {
        val date = DateFormat.format("yyyyMMdd", Calendar.getInstance())
        return "$username($date)"
    }
}
