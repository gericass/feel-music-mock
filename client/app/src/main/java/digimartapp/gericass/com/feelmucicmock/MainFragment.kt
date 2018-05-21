package digimartapp.gericass.com.feelmucicmock


import android.os.Bundle
import android.support.v4.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import digimartapp.gericass.com.feelmucicmock.databinding.FragmentMainBinding


class MainFragment : Fragment() {

    var mViewmodel: MainViewModel? = null

    var mMainNavigator: MainNavigator? = null


    override fun onCreateView(inflater: LayoutInflater?, container: ViewGroup?,
                              savedInstanceState: Bundle?): View? {
        // Inflate the layout for this fragment
        val view = inflater!!.inflate(R.layout.fragment_main, container, false)
        val viewDataBinding = FragmentMainBinding.bind(view)
        viewDataBinding.viewmodel = MainViewModel(mMainNavigator)
        return view
    }

}
