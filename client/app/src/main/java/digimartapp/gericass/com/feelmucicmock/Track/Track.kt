package digimartapp.gericass.com.feelmucicmock.track

import java.io.Serializable

/**
 * Created by keita on 2018/06/11.
 */
data class Track(var name: String,
                 var url: String,
                 var comment: String) : Serializable