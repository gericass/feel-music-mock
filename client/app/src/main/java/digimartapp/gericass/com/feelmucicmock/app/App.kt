package digimartapp.gericass.com.feelmucicmock.app

import android.app.Activity
import android.app.Application
import dagger.android.DaggerApplication
import dagger.android.DispatchingAndroidInjector
import dagger.android.HasActivityInjector
import javax.inject.Inject

/**
 * Created by keita on 2018/05/30.
 */

class App : DaggerApplication() {

    override fun applicationInjector() = DaggerAppComponent.builder()
            .application(this)
            .build()

}