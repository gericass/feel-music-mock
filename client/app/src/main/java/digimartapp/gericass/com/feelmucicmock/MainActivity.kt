package digimartapp.gericass.com.feelmucicmock

import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.widget.LinearLayout
import digimartapp.gericass.com.feelmucicmock.constants.MAINFRAGMENT

class MainActivity : AppCompatActivity(), MainNavigator {

    private lateinit var mMainViewModel: MainViewModel

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val mainFragment = MainFragment()
        mMainViewModel = MainViewModel(this, applicationContext)
        mainFragment.mViewmodel = mMainViewModel


        val transaction = supportFragmentManager.beginTransaction()
        transaction.replace(R.id.main_fragment_container, mainFragment)
        transaction.commit()

    }

    override fun getSensor() {

    }

    override fun onDestroy() {
        mMainViewModel.onActivityDestroyed()
        super.onDestroy()
    }
}
