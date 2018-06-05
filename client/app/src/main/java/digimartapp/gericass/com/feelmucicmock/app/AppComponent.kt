package digimartapp.gericass.com.feelmucicmock.app

import dagger.BindsInstance
import dagger.Component
import dagger.android.AndroidInjectionModule
import dagger.android.AndroidInjector
import javax.inject.Singleton

/**
 * Created by keita on 2018/05/30.
 */

@Singleton
@Component(modules = arrayOf(AndroidInjectionModule::class, AndroidModule::class))
interface AppComponent : AndroidInjector<App> {
    @Component.Builder
    interface Builder {
        @BindsInstance
        fun application(application: App): Builder
        fun build(): AppComponent
    }
    override fun inject(app: App)
}