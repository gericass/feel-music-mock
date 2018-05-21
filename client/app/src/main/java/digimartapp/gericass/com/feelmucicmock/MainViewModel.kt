package digimartapp.gericass.com.feelmucicmock

/**
 * Created by keita on 2018/05/21.
 */

class MainViewModel(private var mMainNavigator: MainNavigator?) {

    fun onActivityDestroyed() {
        mMainNavigator = null
    }

    fun getSensor() {
        if(mMainNavigator != null) {
        }
    }

}