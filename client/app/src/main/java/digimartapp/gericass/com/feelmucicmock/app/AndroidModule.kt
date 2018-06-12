package digimartapp.gericass.com.feelmucicmock.app

import dagger.Module
import dagger.android.ContributesAndroidInjector
import digimartapp.gericass.com.feelmucicmock.MainActivity
import digimartapp.gericass.com.feelmucicmock.MainFragment

/**
 * Created by keita on 2018/05/30.
 */

@Module
abstract class AndroidModule {
    @ContributesAndroidInjector
    abstract fun contributeMainActivity(): MainActivity
    @ContributesAndroidInjector
    abstract fun contributeMainFragment(): MainFragment
}