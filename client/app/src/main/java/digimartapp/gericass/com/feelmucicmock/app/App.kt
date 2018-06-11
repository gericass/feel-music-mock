package digimartapp.gericass.com.feelmucicmock.app

import android.app.Activity
import android.app.Application
import android.app.Fragment
import dagger.android.AndroidInjector
import dagger.android.DaggerApplication
import dagger.android.DispatchingAndroidInjector
import dagger.android.HasActivityInjector
import dagger.android.support.HasSupportFragmentInjector
import javax.inject.Inject

/**
 * Created by keita on 2018/05/30.
 */

class App : DaggerApplication(), HasSupportFragmentInjector {


    override fun applicationInjector() = DaggerAppComponent.builder()
            .application(this)
            .build()

    @Inject
    lateinit var androidInjector: DispatchingAndroidInjector<Fragment>

    @Suppress("UNCHECKED_CAST")
    override fun supportFragmentInjector() = androidInjector as AndroidInjector<android.support.v4.app.Fragment>
}