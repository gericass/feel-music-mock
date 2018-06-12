package digimartapp.gericass.com.feelmucicmock.api

import digimartapp.gericass.com.feelmucicmock.track.Track
import retrofit2.http.GET
import io.reactivex.Observable


/**
 * Created by keita on 2018/06/12.
 */

interface TrackClient {
    @GET("flick/track")
    fun getFlickTrack(): Observable<Track>
}