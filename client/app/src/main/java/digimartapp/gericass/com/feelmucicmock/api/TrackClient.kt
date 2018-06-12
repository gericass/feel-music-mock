package digimartapp.gericass.com.feelmucicmock.api

import digimartapp.gericass.com.feelmucicmock.constants.GET_FLICK_TRACK
import digimartapp.gericass.com.feelmucicmock.track.Track
import retrofit2.http.GET
import io.reactivex.Observable


/**
 * Created by keita on 2018/06/12.
 */

interface TrackClient {
    @GET(GET_FLICK_TRACK)
    fun getFlickTrack(): Observable<Track>
}