package digimartapp.gericass.com.feelmucicmock

import dagger.Module
import dagger.Provides

/**
 * Created by keita on 2018/06/05.
 */

@Module
class MainModule {
    @Provides
    fun provideMainNavigator(mainActivity: MainActivity): MainNavigator = mainActivity

    @Provides
    fun provideMainViewModel(mainNavigator: MainNavigator) = MainViewModel(mainNavigator)
}