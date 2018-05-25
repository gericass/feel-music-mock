package digimartapp.gericass.com.feelmucicmock

import android.content.Context
import android.databinding.BaseObservable
import android.databinding.Bindable

/**
 * Created by keita on 2018/05/21.
 */

class MainViewModel(private var mMainNavigator: MainNavigator?,
                    private val mContext: Context) : BaseObservable() {

    @Bindable
    var sensors: StringBuffer? = null
        set(value) {
            field = value
            notifyPropertyChanged(BR.sensors)
        }

    @Bindable
    var sensorValue: String? = null
        set(value) {
            field = value
            notifyPropertyChanged(BR.sensorValue)
        }

    fun onActivityDestroyed() {
        mMainNavigator = null
    }

    fun onActivityResumed() {
        mMainNavigator?.getActiveSensor()
    }


}