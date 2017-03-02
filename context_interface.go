package swagger

import (
	"net/http"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

/**
 * http.Requestが発生するごとに生成されるContext
 *
 * 内部の具体的実装はGAE/Go Appの実装者に任せられる。
 */
type RequestContext interface {
	/**
	 * Validatorを生成させる
	 *
	 * Contextをチェックし、適したValidatorを生成させる。
	 *
	 * ex) swagger.NewValidator()
	 */
	NewValidatorFactory() ValidatorFactory

	/**
	 * request -> parameterへのバインド制御インターフェースを生成する
	 *
	 * ex) swagger.NewRequestBinder()
	 */
	NewRequestBinder(request *http.Request) RequestBinder

	/**
	 * Request -> Parameterのバインド失敗時に呼び出される。
	 */
	NewBindErrorResponse(err error) middleware.Responder

	/**
	 * ハンドリングの完了処理を行う。
	 *
	 * このメソッドは制御の最後にかならず呼び出される。
	 * 必要に応じてリソースの開放処理を行う。
	 */
	Done(writer http.ResponseWriter, response middleware.Responder)
}

type ContextFactory interface {
	/**
	 * 1ハンドリングごとのコンテキストを生成する
	 */
	NewContext(request *http.Request) RequestContext
}


/**
 * APIエンドポイント(&METHOD) ごとに用意されるハンドリングデータ
 */
type HandleRequest struct {
	// /path/to/api
	Path        string

	// GET, POST, PUT...
	Method      string

	// Function
	// DefaultApiController.
	HandlerFunc func(context RequestContext, request *http.Request) middleware.Responder
}

/**
 * HandleRequestと実際のRouterのマッピングを行なう。
 */
type HandleMapper interface {
	/**
	 * リクエストハンドラを追加する
	 */
	PutHandler(request HandleRequest)

	/**
	 * 最終的なハンドリングを行なうためのRouterを生成する
	 */
	NewRouter(controller ContextFactory) *mux.Router
}
