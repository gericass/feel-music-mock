package digimartapp.gericass.com.feelmucicmock.home

import android.content.Context
import android.databinding.BaseObservable
import android.databinding.Bindable
import digimartapp.gericass.com.feelmucicmock.BR
import digimartapp.gericass.com.feelmucicmock.MainNavigator

/**
 * Created by keita on 2018/05/29.
 */

class HomeViewModel(private var mHomeNavigator: HomeNavigator,
                    private val mContext: Context) : BaseObservable()