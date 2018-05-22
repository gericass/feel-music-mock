package digimartapp.gericass.com.feelmucicmock

import android.content.Context

/**
 * Created by keita on 2018/05/21.
 */

class MainViewModel(private var mMainNavigator: MainNavigator?,
                    private val mContext: Context) {


    fun onActivityDestroyed() {
        mMainNavigator = null
    }

    fun getSensor() {
        if (mMainNavigator != null) {
        }
    }

}