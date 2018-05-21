package digimartapp.gericass.com.feelmucicmock

import android.support.v7.app.AppCompatActivity
import android.os.Bundle

class MainActivity : AppCompatActivity(), MainNavigator {

    private lateinit var mMainViewModel: MainViewModel

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val mainFragment = MainFragment()
        mMainViewModel = MainViewModel(this)
        mainFragment.mViewmodel = mMainViewModel

    }

    override fun getSensor() {

    }

    override fun onDestroy() {
        mMainViewModel.onActivityDestroyed()
        super.onDestroy()
    }
}
