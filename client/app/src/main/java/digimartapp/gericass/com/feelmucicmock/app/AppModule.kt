package digimartapp.gericass.com.feelmucicmock.app

import android.app.Application
import android.content.Context
import dagger.Module
import dagger.Provides
import javax.inject.Singleton

/**
 * Created by keita on 2018/06/12.
 */

@Module
class AppModule {
    @Singleton
    @Provides
    fun provideContext(application: Application): Context = application
}